package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"github.com/spf13/cobra"
	"log"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall CNS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		manager := internal.NewCnsManager()
		executor, err := internal.NewCommandExecutor(manager, []string{})
		if err != nil {
			log.Println(err)
			return
		}

		if err := executor.UninstallCNS(); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
