package goavcodec

import (
	"fmt"
	"strconv"
)

// Options ...
type Options struct {
	Speed *float64
}

// Set ...
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

// SetSpeed ...
func (options *Options) SetSpeed(speed float64) *Options {
	options.Speed = &speed
	return options
}

// ToArgs ...
func (options *Options) ToArgs() []string {
	args := []string{}
	if options.Speed != nil && *options.Speed != 1 {
		args = append(args, "-vf", fmt.Sprintf("setpts=(1/%v)*PTS", *options.Speed))
	}
	return args
}
