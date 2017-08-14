package goavcodec

import (
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestConvert(t *testing.T) {

	src, err := os.Open("./testdata/src/origin.webm")
	Expect(t, err).ToBe(nil)
	Expect(t, src).TypeOf("*os.File")
	defer src.Close()

	dest, err := os.Create("./testdata/dest/dest.mp4")
	Expect(t, err).ToBe(nil)
	Expect(t, dest).TypeOf("*os.File")
	defer func() {
		dest.Close()
		os.Remove(dest.Name())
	}()

	encoder, err := NewEncoder(src)
	Expect(t, err).ToBe(nil)
	Expect(t, encoder).TypeOf("*goavcodec.Encoder")
	err = encoder.Encode(dest)
	Expect(t, err).ToBe(nil)
	// sd, err := dest.Stat()
	// Expect(t, err).ToBe(nil)
	// Expect(t, sd.Size()).Not().ToBe(int64(0))
}
