package goavcodec

// #cgo pkg-config: libavformat libavcodec
// #include <libavformat/avformat.h>
// #include <libavcodec/avcodec.h>
import "C"

import (
	"os"
	"path/filepath"
)

// Encoder ...
type Encoder struct {
	Context *C.struct_AVFormatContext
	buf     []byte
}

// NewEncoder ...
func NewEncoder(src *os.File) (*Encoder, error) {

	C.av_register_all()

	fullpath, err := filepath.Abs(src.Name())
	if err != nil {
		return nil, err
	}
	var ctx *C.struct_AVFormatContext
	ret := C.avformat_open_input(
		&ctx,
		C.CString(fullpath),
		nil, nil,
	)
	if err := Error(ret); err != nil {
		return nil, err
	}
	return &Encoder{
		Context: ctx,
	}, nil
}

// Encode ...
func (encoder *Encoder) Encode(out *os.File) error {
	return nil
}
