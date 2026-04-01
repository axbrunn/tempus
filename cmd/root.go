package cmd

import (
	"fmt"
	"os"

	"github.com/axbrunn/tempus/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tempus",
	Short: "Comptime registratie TUI",
	Long:  "Een TUI app voor het bijhouden van vrijedagen en overuren.",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
