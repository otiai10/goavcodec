package goavcodec

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/otiai10/mint"
)

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

	os.Remove(dest)
}
