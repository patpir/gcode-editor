package gcode

import (
	"github.com/256dpi/gcode"
)

type Params map[string]float64

func lineParams(line gcode.Line) Params {
	params := make(Params)
	for _, code := range line.Codes[1:] {
		if code.Letter != "" {
			params[code.Letter] = code.Value
		}
	}
	return params
}
