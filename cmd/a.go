package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"github.com/spf13/cobra"
	"log"
)

var aCmd = &cobra.Command{
	Use:                "a",
	Short:              "Run a command",
	Long:               ``,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		manager := internal.NewCnsManager()
		var flag int
		for idx, val := range args {
			if val == "-i" {
				flag = idx + 1
			}
		}
		var executor *internal.CommandExecutor
		var err error
		if flag != 0 {
			executor, err = internal.NewCommandExecutor(manager, args[flag+1:])
		} else {
			executor, err = internal.NewCommandExecutor(manager, args)
		}

		if err != nil {
			log.Println(err)
			return
		}

		if err := executor.Execute(args[flag]); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(aCmd)
}
