package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/patpir/gcode-pattern-detect/gcode"
	"github.com/patpir/gcode-pattern-detect/pattern"
	"github.com/patpir/svgshapes"
)

func main() {
	signal.Ignore(syscall.SIGPIPE)

	file, err := os.Open("example.gcode")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	interpreter, err := gcode.Interpret(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error interpreting file: %v\n", err)
		os.Exit(1)
	}

/*
	commentCount := 0
	interpreter.Handler = gcode.Handler{
		HandleMove: func(move gcode.MoveCommand) {
			fmt.Printf("Position after move: X=%f Y=%f Z=%f E=%f F=%f\n",
				move.X,
				move.Y,
				move.Z,
				move.E,
				move.FeedRate,
			)
		},
		HandleUnknown: func(cmd gcode.UnknownCommand) {
			fmt.Printf("UNKNOWN COMMAND: %v", cmd)
		},
		HandleComment: func(string) {
			commentCount += 1
		},
	}
	if err := interpreter.Walk(); err != nil {
		fmt.Fprintf(os.Stderr, "Error in GCode: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Suppressed output of %d comments\n", commentCount)
*/

	detector := pattern.Detector{
		Interpreter: interpreter,
	}
	if err := detector.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while loading data into detector: %v\n", err)
		os.Exit(1)
	}

	channels := detector.Channels
	channels.DeriveAll()

	svgDoc := svgshapes.NewDocument("")
	pattern.Plot(svgDoc, channels, []pattern.PlotInfo{
		{
			Id: "X",
			Name: "X position",
			Color: "#00aa00",
		},
		{
			Id: "Y",
			Name: "Y position",
			Color: "#aa0000",
		},
		{
			Id: "Z",
			Name: "Z position",
			Color: "#000000",
			Scaling: 10,
		},
		{
			Id: "E'",
			Name: "Relative extrusion",
			Color: "#0000aa",
			Scaling: 50,
		},
		{
			Id: "FeedRate",
			Name: "Speed",
			Color: "#888888",
			Scaling: 0.10,
		},
	})
	outFile, err := os.Create("output.svg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()
	svgDoc.WriteIndent(outFile)
}
