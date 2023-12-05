/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"restiq/helper"

	"github.com/spf13/cobra"
)

var backupAllCmd = &cobra.Command{
	Use:   "backupall",
	Short: "Backup All Repos",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		helper.ReadViperConfig()
		helper.ResticBackupRepoNew(args[0])
	},
}

func init() {
	rootCmd.AddCommand(backupAllCmd)
}
