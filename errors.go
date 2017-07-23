package goavcodec

import "C"

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

// "errors" represents all errors from ffmpeg
// @see http://ffmpeg.org/doxygen/trunk/error_8h_source.html for more information
var errors = []*Err{
	&Err{
		-mktag('I', 'N', 'D', 'A'),
		"Invalid data found when processing input",
	},
}

// Error provide specific error for code.
func Error(code C.int) error {
	c := int32(code)
	for _, err := range errors {
		if err.Code == c {
			return err
		}
	}
	return nil
}
