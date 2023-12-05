/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"restiq/helper"

	"github.com/spf13/cobra"
)

var snapshotsCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "Lists Snapshots for Repo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		helper.ReadViperConfig()
		helper.ResticSnapshots(args)
	},
}

func init() {
	rootCmd.AddCommand(snapshotsCmd)
}
