package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := getRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "gcpd",
		Short: "gcode pattern detector",
		Long: "Detect and modify patterns in GCode files",
	}
	rootCmd.AddCommand(getVisualizeCmd())
	return rootCmd
}
