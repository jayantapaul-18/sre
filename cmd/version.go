package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "development-v0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of SRE",
	Long:  `All software has versions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:\t", Version)
	},
}
