package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "croctransform",
	Short: "A tool for analyzing fossil data to identify reptile locomotion patterns",
	Long: `CrocTransform is a scientific data analysis tool that helps paleontologists
and researchers identify unusual locomotion transitions in ancient reptiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}