/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current session",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		osType := runtime.GOOS
		homeDir, _ := os.UserHomeDir()
		executor := internal.NewCommandExecutor(osType, homeDir, []string{})
		executor.DestroySession()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
