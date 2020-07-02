package gcode

type HeadPosition struct {
	X float64
	Y float64
	Z float64
}

type Position struct {
	HeadPosition
	E float64
}
