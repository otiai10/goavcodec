package goavcodec

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	. "github.com/otiai10/mint"
)

func TestMain(m *testing.M) {
	os.MkdirAll("../../testdata/dest", os.ModePerm)
	code := m.Run()
	os.RemoveAll("../../testdata/dest")
	os.Exit(code)
}

func srcpath(name string) string {
	src, err := filepath.Abs(fmt.Sprintf("../../testdata/src/%s", name))
	if err != nil {
		panic(err)
	}
	return src
}

func destpath(name string) string {
	ext := filepath.Ext(name)
	label := strings.Replace(filepath.Base(name), ext, "", 1)
	dest, err := filepath.Abs(fmt.Sprintf("../../testdata/dest/%s.%v%s", label, time.Now().Unix(), ext))
	if err != nil {
		panic(err)
	}
	return dest
}

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	Expect(t, err).ToBe(nil)
	Expect(t, client).TypeOf("*goavcodec.Client")

	Because(t, "NewClient looks for ffmpeg binary from PATH", func(t *testing.T) {
		envpath := os.Getenv("PATH")
		os.Setenv("PATH", "")
		_, err = NewClient()
		Expect(t, err).Not().ToBe(nil)
		os.Setenv("PATH", envpath)
	})

	When(t, "Specified bin path is not found", func(t *testing.T) {
		_, err := NewClient("/foobar")
		Expect(t, err).Not().ToBe(nil)
	})
	When(t, "Specified bin path is not executable", func(t *testing.T) {
		_, err := NewClient("./all_test.go")
		Expect(t, err).Not().ToBe(nil)
	})
}

func TestClient_Convert(t *testing.T) {

	client, err := NewClient()
	Expect(t, err).ToBe(nil)
	Expect(t, client).TypeOf("*goavcodec.Client")

	src := srcpath("origin.webm")
	dest := destpath("normal.mp4")
	Expect(t, err).ToBe(nil)

	err = client.Convert(src, dest)
	Expect(t, err).ToBe(nil)

	info, err := os.Stat(dest)
	Expect(t, err).ToBe(nil)
	Expect(t, info.IsDir()).ToBe(false)
	Expect(t, info.Size() > 100).ToBe(true)

	When(t, "not-existing file path is given", func(t *testing.T) {
		client, _ := NewClient()
		src := srcpath("notfound.webm")
		dest := destpath("foo.mp4")
		err = client.Convert(src, dest)
		Expect(t, err).Not().ToBe(nil)
		Expect(t, bytes.HasSuffix(client.StdErr, []byte("No such file or directory\n"))).ToBe(true)
	})

	When(t, "Options specified", func(t *testing.T) {
		client, err := NewClient()
		Expect(t, err).ToBe(nil)
		src := srcpath("origin.webm")
		Expect(t, err).ToBe(nil)
		dest := destpath("double_speed.mp4")
		Expect(t, err).ToBe(nil)
		opt := new(Options)
		opt.Set("speed", 2).Set("speed", 2.0).Set("speed", "2")
		err = client.Convert(src, dest, opt)
		Expect(t, err).ToBe(nil)
	})
}
