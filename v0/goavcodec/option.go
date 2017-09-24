package goavcodec

import (
	"fmt"
	"strconv"
	"time"
)

// MakeArgFunc represents function which can make argument of avconv binary.
type MakeArgFunc func() []string

// Options ...
type Options struct {
	Speed    MakeArgFunc
	Start    MakeArgFunc
	Duration MakeArgFunc
}

// Set is an alias for "SetXXX" methods.
func (options *Options) Set(name string, val interface{}) *Options {
	switch name {
	case "speed":
		if speed, ok := val.(float64); ok {
			options.SetSpeed(speed)
		} else if speed, ok := val.(int); ok {
			options.SetSpeed(float64(speed))
		} else if speed, err := strconv.ParseFloat(fmt.Sprintf("%v", val), 32); err == nil {
			options.SetSpeed(speed)
		}
	case "start":
		if start, ok := val.(string); ok {
			if dur, err := time.ParseDuration(start); err == nil {
				options.SetStart(dur)
			}
		}
	case "duration":
		if duration, ok := val.(string); ok {
			if dur, err := time.ParseDuration(duration); err == nil {
				options.SetDuration(dur)
			}
		}
	}
	return options
}

// SetSpeed sets speed option.
func (options *Options) SetSpeed(speed float64) *Options {
	if speed == 1 {
		return options
	}
	options.Speed = func() []string {
		return []string{"-vf", fmt.Sprintf("setpts=(1/%v)*PTS", speed)}
	}
	return options
}

// SetStart sets start time.
func (options *Options) SetStart(start time.Duration) *Options {
	if start == 0 {
		return options
	}
	h := int(start.Hours())
	start = start - time.Duration(h)*time.Hour
	m := int(start.Minutes())
	start = start - time.Duration(m)*time.Minute
	s := int(start.Seconds())
	options.Start = func() []string {
		return []string{
			"-ss",
			fmt.Sprintf("%02d:%02d:%02d", h, m, s),
		}
	}
	return options
}

// SetDuration sets end time.
func (options *Options) SetDuration(dur time.Duration) *Options {
	if dur == 0 {
		return options
	}
	h := int(dur.Hours())
	dur = dur - time.Duration(h)*time.Hour
	m := int(dur.Minutes())
	dur = dur - time.Duration(m)*time.Minute
	s := int(dur.Seconds())
	options.Duration = func() []string {
		return []string{
			"-t",
			fmt.Sprintf("%02d:%02d:%02d", h, m, s),
		}
	}
	return options
}
