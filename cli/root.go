package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

/* Constants */
const VERSION = "0.1"

/* Flags */
var Verbose bool
var ConfigPath string

var rootCmd = &cobra.Command{
	Use:     "dizzy",
	Version: VERSION,
	Short:   "Dizzy is a powerful CLI tool to quickly visualize data.",
	Long:    `Display 'dizzy' Fuzzy is a tool that helps you visualize and manage command outputs in an elegant way.`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello dizzy !")
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display logs.")
	rootCmd.PersistentFlags().StringVar(&ConfigPath, "config", "", "Use a specific config.")

	rootCmd.Execute()
}
