/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version info",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		bi, _ := debug.ReadBuildInfo()
		fmt.Printf("+%v\n", bi)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
