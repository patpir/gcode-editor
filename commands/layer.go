package commands

import (
	"github.com/spf13/cobra"
)

func getLayerCmd() *cobra.Command {
	layerCmd := &cobra.Command{
		Use: "layer",
		Short: "commands to analyse and modify layers",
	}
	layerCmd.AddCommand(
		//getLayerRemoveCmd(),
		//getLayerCopyCmd(),
		getLayerListCmd(),
	)
	return layerCmd
}
