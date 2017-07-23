package goavcodec

// #cgo LDFLAGS: -L /usr/local/lib -lavcodec
// #include <libavcodec/avcodec.h>
// #include <libavformat/avformat.h>
import "C"

func init() {
	C.avcodec_register_all()
}

type (
	Dictionary      C.struct_AVDictionary
	InputFormat     C.struct_AVInputFormat
	AVFormatContext C.struct_AVFormatContext
)
