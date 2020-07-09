package gcode

import (
	"io"

	"github.com/256dpi/gcode"
)

type Layer struct {
	Position     float64
	FileOffset   int
	Instructions []gcode.Line
	Partitioned  bool
	Extrusion    float64
}

func ReadLayers(reader io.Reader) ([]*Layer, error) {
	in, err := Interpret(reader)
	if err != nil {
		return nil, err
	}
	var layers []*Layer
	var currentLayer *Layer
	var extrusionStart float64
	for line, err := in.nextLine(); err != io.EOF; line, err = in.nextLine() {
		if err != nil {
			return layers, err
		}

		currentZOffset := in.Position().Z
		if currentLayer == nil || currentLayer.Position != currentZOffset {
			// apply extrusion to completed layer
			calculateLayerExtrusion(currentLayer, extrusionStart, in.Position().E)

			// find or create layer with this Position as new currentLayer
			currentLayer = nil
			for _, layer := range layers {
				if layer.Position == currentZOffset {
					if layer.Extrusion == 0.0 {
						layer.FileOffset = in.currentLine
						layer.Instructions = nil
					} else {
						layer.Partitioned = true
					}
					currentLayer = layer
				}
			}
			if currentLayer == nil {
				currentLayer = &Layer{
					Position:     currentZOffset,
					FileOffset:   in.currentLine,
				}
				layers = append(layers, currentLayer)
			}
			extrusionStart = in.Position().E
		}

		// add instruction to this layer
		currentLayer.Instructions = append(currentLayer.Instructions, line)
	}
	calculateLayerExtrusion(currentLayer, extrusionStart, in.Position().E)

	return layers, nil
}

func calculateLayerExtrusion(layer *Layer, start float64, end float64) {
	if layer != nil && start != end {
		layer.Extrusion += end - start
	}
}
