package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"github.com/spf13/cobra"
	"log"
)

// eCmd represents the e command
var eCmd = &cobra.Command{
	Use:   "e",
	Short: "Execute a command by id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		manager := internal.NewCnsManager()
		executor, err := internal.NewCommandExecutor(manager, []string{})
		if err != nil {
			log.Println(err)
			return
		}

		if err := executor.GetCommandById(args[0]).Execute("execution_of_e"); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(eCmd)
}
