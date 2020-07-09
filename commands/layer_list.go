package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/patpir/gcode-editor/gcode"
)

func getLayerListCmd() *cobra.Command {
	var verbose bool
	layerListCmd := &cobra.Command{
		Use: "list <input-file>",
		Short: "show all layers in the file",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runLayerList(cmd, args[0], verbose)
		},
	}
	layerListCmd.Flags().BoolVarP(&verbose, "verbose", "v", false,
		"Print detailed information for each layer")
	return layerListCmd
}

func runLayerList(cmd *cobra.Command, inputFileName string, verbose bool) {
	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	layers, err := gcode.ReadLayers(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error interpreting file: %v\n", err)
		os.Exit(1)
	}

	if verbose {
		fmt.Println("Layer   Extrusion   Line  Instr")
	}
	for _, layer := range layers {
		if verbose {
			fmt.Printf("%5.1f  %7.2f mm   %4d  %5d\n",
				layer.Position,
				layer.Extrusion,
				layer.FileOffset,
				len(layer.Instructions),
			)
		} else {
			fmt.Println(layer.Position)
		}
		if layer.Partitioned {
			fmt.Fprintf(os.Stderr, "Layer %f is partitioned\n", layer.Position)
		}
	}
}
