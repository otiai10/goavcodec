package goavcodec

import (
	"fmt"
	"strconv"
)

// MakeArgFunc represents function which can make argument of avconv binary.
type MakeArgFunc func() []string

// Options ...
type Options struct {
	Speed MakeArgFunc
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

// ToArgs ...
func (options *Options) ToArgs() []string {
	args := []string{}
	if options.Speed != nil {
		args = append(args, options.Speed()...)
	}
	return args
}
