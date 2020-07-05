package commands

import (
	"fmt"
	"os"

	"github.com/patpir/gcode-pattern-detect/gcode"
	"github.com/patpir/gcode-pattern-detect/pattern"
	"github.com/patpir/svgshapes"
	"github.com/spf13/cobra"
)

func getVisualizeCmd() *cobra.Command {
	var outputFileName string
	visualizeCmd := &cobra.Command{
		Use: "visualize <input-file>",
		Short: "create a SVG visualization",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runVisualize(cmd, args[0], outputFileName)
		},
	}
	visualizeCmd.Flags().StringVarP(&outputFileName, "output", "o", "output.svg",
		"File path to output SVG file")
	return visualizeCmd
}

func runVisualize(cmd *cobra.Command, inputFileName, outputFileName string) {
	file, err := os.Open(inputFileName)
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
	outFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()
	svgDoc.WriteIndent(outFile)
}
