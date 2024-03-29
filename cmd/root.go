package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cns",
	Short: "CLI NameSpace (CNS) is a GoLang-based CLI application designed to manage command sessions for users.",
	Long: `CLI NameSpace (CNS) is a GoLang-based CLI application designed to manage command sessions for users. 
It allows users to store, retrieve, and execute commands within named sessions, enhancing the command-line experience by organizing commands in a user-friendly manner. 
CNS is perfect for users who frequently work with long or complex commands and want to save them for future use.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
