package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all commands inside this session",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		homedir, _ := os.UserHomeDir()
		lister := internal.NewLister(homedir)
		lister.ListCommands()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
