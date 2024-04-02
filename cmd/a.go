package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:                "a",
	Short:              "Run a command",
	Long:               ``,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		var flag int
		for idx, val := range args {
			if val == "-i" {
				flag = idx + 1
			}
		}
		var executor *internal.CommandExecutor
		if flag != 0 {
			executor = internal.NewCommandExecutor(runtime.GOOS, homeDir, args[flag+1:])
		} else {
			executor = internal.NewCommandExecutor(runtime.GOOS, homeDir, args)
		}

		executor.Execute(args[flag])
	},
}

func init() {
	rootCmd.AddCommand(aCmd)
}
