package goavcodec

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
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

	src, err := filepath.Abs("../../testdata/src/origin.webm")
	Expect(t, err).ToBe(nil)
	dest, err := filepath.Abs(fmt.Sprintf("../../testdata/dest/dest.%v.mp4", time.Now().Unix()))
	Expect(t, err).ToBe(nil)

	err = client.Convert(src, dest)
	Expect(t, err).ToBe(nil)

	info, err := os.Stat(dest)
	Expect(t, err).ToBe(nil)
	Expect(t, info.IsDir()).ToBe(false)
	Expect(t, info.Size() > 100).ToBe(true)

	When(t, "not-existing file path is given", func(t *testing.T) {
		client, _ := NewClient()
		src, _ := filepath.Abs("notfound.webm")
		dest, _ := filepath.Abs(fmt.Sprintf("../../testdata/dest/%v.mp4", time.Now().Unix()))
		err = client.Convert(src, dest)
		Expect(t, err).Not().ToBe(nil)
		Expect(t, bytes.HasSuffix(client.StdErr, []byte("No such file or directory\n"))).ToBe(true)
	})
}
