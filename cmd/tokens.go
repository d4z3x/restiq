/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"restiq/helper"
	"strings"

	"github.com/spf13/cobra"
)

var tokensCmd = &cobra.Command{
	Use:   "tokens",
	Short: "List tokens",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		helper.ReadViperConfig()
		for k, v := range helper.C.Token {
			fmt.Printf("%s=%s%s\n", k, v[0:4], strings.Repeat("*", len(v)-5)[4:])
		}
	},
}

func init() {
	rootCmd.AddCommand(tokensCmd)
}
