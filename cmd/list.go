package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"github.com/spf13/cobra"
	"log"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all commands inside this session",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		manager := internal.NewCnsManager()
		lister, err := internal.NewLister(manager)
		if err != nil {
			log.Println(err)
			return
		}
		lister.ListCommands()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
