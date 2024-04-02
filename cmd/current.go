/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/AdCodeLabs/cns/internal"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current session",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		osType := runtime.GOOS
		homeDir, _ := os.UserHomeDir()
		executor := internal.NewCommandExecutor(osType, homeDir, []string{})
		fmt.Println(executor.GetCurrentSession())
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
