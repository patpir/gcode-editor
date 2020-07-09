package pattern

import (
	"github.com/patpir/gcode-editor/analysis"
	"github.com/patpir/gcode-editor/gcode"
)

type Detector struct {
	Interpreter *gcode.Interpreter
	Channels    analysis.SourceGroup
}

func (d *Detector) Load() error {
	d.Channels = make(analysis.SourceGroup)
	d.Channels["X"] = &analysis.Source{ ValueUnit: "mm" }
	d.Channels["Y"] = &analysis.Source{ ValueUnit: "mm" }
	d.Channels["Z"] = &analysis.Source{ ValueUnit: "mm" }
	d.Channels["E"] = &analysis.Source{ ValueUnit: "mm" }
	d.Channels["FeedRate"] = &analysis.Source{ ValueUnit: "mm/min" }
	step := 0
	d.Interpreter.Handler = gcode.Handler{
		HandleMove: func(move gcode.MoveCommand) {
			d.Channels["X"].Add(float64(step), move.X)
			d.Channels["Y"].Add(float64(step), move.Y)
			d.Channels["Z"].Add(float64(step), move.Z)
			d.Channels["E"].Add(float64(step), move.E)
			d.Channels["FeedRate"].Add(float64(step), move.FeedRate)
			step += 1
		},
	}
	return d.Interpreter.Walk()
}
