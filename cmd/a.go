package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Run a command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		executor := internal.NewCommandExecutor(runtime.GOOS, homeDir, args)
		executor.Execute()
	},
}

func init() {
	rootCmd.AddCommand(aCmd)
}
