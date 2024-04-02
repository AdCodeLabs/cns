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

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall CNS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uninstall called")
		osType := runtime.GOOS
		homeDir, _ := os.UserHomeDir()
		executor := internal.NewCommandExecutor(osType, homeDir, []string{})
		executor.UninstallCNS()
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
