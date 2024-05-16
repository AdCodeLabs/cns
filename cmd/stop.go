/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"github.com/spf13/cobra"
	"log"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current session",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		manager := internal.NewCnsManager()
		executor, _ := internal.NewCommandExecutor(manager, []string{})
		executor.DestroySession()
		log.Println("Destroyed current session...")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
