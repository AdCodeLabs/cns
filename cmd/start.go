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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		start := internal.NewStarter()
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
