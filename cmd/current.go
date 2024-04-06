/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"github.com/spf13/cobra"
	"log"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current session",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		manager := internal.NewCnsManager()
		executor, err := internal.NewCommandExecutor(manager, []string{})
		if err != nil {
			log.Println(err)
			return
		}

		sess, err := executor.GetCurrentSession()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Your current session is: ", sess)
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
