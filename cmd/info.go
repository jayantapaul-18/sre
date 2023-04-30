package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Project info",
	Long:  `Project details info will be in github wiki`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info sections")
	},
}
