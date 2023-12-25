package cmd

import (
	"restiq/helper"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists Repos",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		helper.ReadViperConfig()
		helper.ResticListRepos()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
