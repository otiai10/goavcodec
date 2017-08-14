package goavcodec

import "C"
import "fmt"

// ftp://ftp.acer.com.tr/Archiv/gpl/AS9100/GPL_AS9100/MPlayer-0.90/libavcodec/avcodec.c
func mktag(a, b, c, d int) int32 {
	return int32(a | (b << 8) | (c << 16) | (d << 24))
}

// Err is an alias for FFmpeg Errors
type Err struct {
	Code    int32
	Message string
}

// Error to satisfy `error` interface
func (e *Err) Error() string {
	return e.Message
}

// EOF represents "End of file" error, to specify end of read loop
var EOF = &Err{-mktag('E', 'O', 'F', ' '), "End of file"}

// "errors" represents all errors from ffmpeg
// @see http://ffmpeg.org/doxygen/trunk/error_8h_source.html for more information
var errors = []*Err{
	&Err{-mktag(0xF8, 'B', 'S', 'F'), "Bitstream filter not found"},
	&Err{-mktag('B', 'U', 'G', '!'), "Internal bug, also see AVERROR_BUG2"},
	&Err{-mktag('B', 'U', 'F', 'S'), "Buffer too small"},
	&Err{-mktag(0xF8, 'D', 'E', 'C'), "Decoder not found"},
	&Err{-mktag(0xF8, 'D', 'E', 'M'), "Demuxer not found"},
	&Err{-mktag(0xF8, 'E', 'N', 'C'), "Encoder not found"},
	EOF,
	&Err{-mktag('E', 'X', 'I', 'T'), "Immediate exit was requested; the called function should not be restarted"},
	&Err{-mktag('E', 'X', 'T', ' '), "Generic error in an external library"},
	&Err{-mktag(0xF8, 'F', 'I', 'L'), "Filter not found"},
	&Err{-mktag('I', 'N', 'D', 'A'), "Invalid data found when processing input"},
	&Err{-mktag(0xF8, 'M', 'U', 'X'), "Muxer not found"},
	&Err{-mktag(0xF8, 'O', 'P', 'T'), "Option not found"},
	&Err{-mktag('P', 'A', 'W', 'E'), "Not yet implemented in FFmpeg, patches welcome"},
	&Err{-mktag(0xF8, 'P', 'R', 'O'), "Protocol not found"},
	&Err{-mktag(0xF8, 'S', 'T', 'R'), "Stream not found"},
	&Err{-mktag('B', 'U', 'G', ' '), "AVError BUG"},
	&Err{-mktag('U', 'N', 'K', 'N'), "Unknown error, typically from an external library"},
	&Err{-mktag(0xF8, '4', '0', '0'), "AVERROR_HTTP_BAD_REQUEST"},
	&Err{-mktag(0xF8, '4', '0', '1'), "AVERROR_HTTP_UNAUTHORIZED"},
	&Err{-mktag(0xF8, '4', '0', '3'), "AVERROR_HTTP_FORBIDDEN"},
	&Err{-mktag(0xF8, '4', '0', '4'), "AVERROR_HTTP_NOT_FOUND"},
	&Err{-mktag(0xF8, '4', 'X', 'X'), "AVERROR_HTTP_OTHER_4XX"},
	&Err{-mktag(0xF8, '5', 'X', 'X'), "AVERROR_HTTP_SERVER_ERROR"},
}

// Error provide specific error for code.
func Error(code C.int) error {
	c := int32(code)
	for _, err := range errors {
		if err.Code == c {
			return err
		}
	}
	if c != 0 {
		return &Err{Code: c, Message: fmt.Sprintf("Unhandled Error Code: %v", c)}
	}
	return nil
}
