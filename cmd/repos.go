/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"restiq/helper"

	"github.com/spf13/cobra"
)

var reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Lists Repos",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		helper.ReadViperConfig()
		helper.ResticListRepos()
	},
}

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
	rootCmd.AddCommand(reposCmd)
	rootCmd.AddCommand(listCmd)
}
