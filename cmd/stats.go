/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"restiq/helper"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Stats on a repo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		helper.ReadViperConfig()
		helper.ResticStatsRepo(args[0])
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
