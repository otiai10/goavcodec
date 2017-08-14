package goavcodec

// #cgo pkg-config: libavformat libavcodec
// #include <libavformat/avformat.h>
// #include <libavcodec/avcodec.h>
//
// AVStream* get_stream(AVStream **streams, int pos)
// {
//   return streams[pos];
// }
//
// AVFormatContext* out_context(char* name)
// {
//   AVFormatContext *out = avformat_alloc_context();
//   snprintf(out->filename, sizeof(out->filename), "%s", name);
//   return out;
// }
//
// AVOutputFormat* out_format(AVFormatContext* out)
// {
//   return out->oformat;
// }
import "C"

import (
	"fmt"
	"os"
	"path/filepath"
)

// Encoder ...
type Encoder struct {
	Context *C.struct_AVFormatContext
	buf     []byte
	src     *os.File
	dst     *os.File
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

	ret = C.avformat_find_stream_info(ctx, nil)
	if err := Error(ret); err != nil {
		return nil, err
	}

	return &Encoder{
		Context: ctx,
		src:     src,
	}, nil
}

// Encode ...
func (encoder *Encoder) Encode(out *os.File) error {

	outpath, err := filepath.Abs(out.Name())
	if err != nil {
		return err
	}
	// C.av_dump_format(encoder.Context, C.int(0), C.CString(outpath), C.int(0))

	index, codec := encoder.findIndex()

	var packet C.struct_AVPacket
	var finished C.int
	frame := C.av_frame_alloc()
	// var stream C.struct_AVStream

	outctx := C.out_context(C.CString(outpath))
	// format := C.out_format(outctx)
	fmt.Printf("%T\n", outctx)
	// fmt.Println("formatがnilか", format, out.Name())

	for C.av_read_frame(encoder.Context, &packet) >= 0 {
		if index == packet.stream_index {
			C.avcodec_decode_video2(codec, frame, &finished, &packet)
		}
		C.av_packet_unref(&packet)
	}
	return nil
}

func (encoder *Encoder) findIndex() (C.int, *C.struct_AVCodecContext) {
	var stream *C.struct_AVStream
	for i := uint(0); i < uint(encoder.Context.nb_streams); i++ {
		stream = C.get_stream(encoder.Context.streams, C.int(i))
		if stream.codec.codec_type == C.AVMEDIA_TYPE_VIDEO {
			return C.int(i), stream.codec
		}
	}
	return 0, stream.codec
}
