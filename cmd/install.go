package cmd

import (
	"github.com/AdCodeLabs/cns/internal"
	"github.com/spf13/cobra"
	"log"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Do `cns install` to use cns on your machine",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Installing CNS on your System...")

		manager := internal.NewCnsManager()
		installer := internal.NewInstaller(manager)

		if err := installer.Install(); err != nil {
			log.Println(err)
		}

		log.Println("Installation Done...")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
