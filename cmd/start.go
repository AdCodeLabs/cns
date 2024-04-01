package cmd

import (
	"fmt"
	"github.com/AdCodeLabs/cns/internal"
	"log"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new session",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		start := internal.NewStarter(args)
		if len(args) == 1 {
			if err := start.StartNewSession(); err != nil {
				log.Println(err)
			}
		} else if len(args) == 0 {
			if err := start.ListSessions(); err != nil {
				log.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
