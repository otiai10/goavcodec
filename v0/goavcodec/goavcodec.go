package goavcodec

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	avconv = "avconv"
)

// Client ...
type Client struct {
	bin     string
	StdOut  []byte
	StdErr  []byte
	Options *Options
}

// NewClient looks path for `ffmpeg` and returns initialized Client.
func NewClient(binpath ...string) (*Client, error) {
	if len(binpath) != 0 {
		info, err := os.Stat(binpath[0])
		if err != nil {
			return nil, err
		}
		if info.Mode()%2 == 0 {
			return nil, fmt.Errorf("path specified with `%s` is not executable", binpath[0])
		}
		return &Client{bin: binpath[0]}, nil
	}
	bin, err := exec.LookPath(avconv)
	if err != nil {
		return nil, fmt.Errorf("failed to find path to binary: %s", err.Error())
	}
	return &Client{bin: bin}, nil
}

// Convert just converts src to dest with using `ffmpeg -i`
func (c *Client) Convert(src, dest string, opts ...*Options) error {
	if len(opts) != 0 {
		c.Options = opts[0]
	}
	args := c.Args(src, dest)
	cmd := exec.Command(c.bin, args...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to pipe stderr: %s", err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to pipe stdout: %s", err.Error())
	}
	errored := make(chan error)
	completed := make(chan bool)

	go func() {
		if c.StdErr, err = ioutil.ReadAll(stderr); err != nil {
			errored <- fmt.Errorf("failed to drain all stderr: %s", err.Error())
		}
		if c.StdOut, err = ioutil.ReadAll(stdout); err != nil {
			errored <- fmt.Errorf("failed to drain all stdout: %s", err.Error())
		}
		completed <- true
		close(errored)
		close(completed)
	}()

	if err = cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command specified with `%s`: %s", c.bin, err.Error())
	}
	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("command has not completed successfully: %s: %s: %v", err.Error(), string(c.StdErr), args)
	}

	select {
	case err := <-errored:
		return err
	case <-completed:
		return nil
	}

}

// Args ...
func (c *Client) Args(src, dest string) []string {
	args := []string{}
	if c.Options != nil && c.Options.Start != nil {
		args = append(args, c.Options.Start()...)
	}
	args = append(args, "-y", "-i", src)
	if c.Options != nil && c.Options.Speed != nil {
		args = append(args, c.Options.Speed()...)
	}
	if c.Options != nil && c.Options.Duration != nil {
		args = append(args, c.Options.Duration()...)
	}
	args = append(args, dest)
	return args
}
