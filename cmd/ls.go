/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"restiq/helper"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files from Snapshot",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		helper.ReadViperConfig()
		helper.ResticLsSnapshot(args)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
