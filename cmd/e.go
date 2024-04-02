package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// eCmd represents the e command
var eCmd = &cobra.Command{
	Use:   "e",
	Short: "Execute a command by id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		osType := runtime.GOOS
		homeDir, _ := os.UserHomeDir()
		executor := internal.NewCommandExecutor(osType, homeDir, []string{})
		executor.GetCommandById(args[0]).Execute()
	},
}

func init() {
	rootCmd.AddCommand(eCmd)
}
