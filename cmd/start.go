package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"log"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new session",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		manager := internal.NewCnsManager()
		start, err := internal.NewStarter(manager, args)
		if err != nil {
			log.Println(err)
		}

		if len(args) == 1 {
			log.Println("Trying to start a new session...")
			if err := start.StartNewSession(); err != nil {
				log.Println(err)
			}
		} else if len(args) == 0 {
			log.Println("Listing available sessions...")
			if err := start.ListSessions(); err != nil {
				log.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
