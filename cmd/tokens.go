/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"restiq/helper"

	"github.com/spf13/cobra"
)

var tokensCmd = &cobra.Command{
	Use:   "tokens",
	Short: "List tokens",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		helper.ReadViperConfig()
		for k, v := range helper.C.Token {
			fmt.Printf("%s=%s\n", k, v)
		}
		// fmt.Printf("Tokens=%+q\n", helper.C.Token)
		// fmt.Printf("Tokens=%+q\n", helper.C.Token)
	},
}

func init() {
	rootCmd.AddCommand(tokensCmd)
}
