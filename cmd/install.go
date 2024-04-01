package cmd

import (
	"fmt"
	"github.com/AdCodeLabs/cns/internal"
	"log"
	"runtime"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
		os := runtime.GOOS
		fmt.Println(os)
		installer := internal.NewInstaller(os)
		if err := installer.Install(); err != nil {
			log.Println(err)
		}
		fmt.Println("installation done")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
