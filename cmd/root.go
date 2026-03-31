package cmd

import (
	"fmt"
	"os"

	"github.com/axbrunn/tempus/internal/store"
	"github.com/axbrunn/tempus/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tempus",
	Short: "Comptime registratie TUI",
	Long:  "Een TUI app voor het bijhouden van vrijedagen en overuren.",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := store.DataPath()
		if err != nil {
			fmt.Println("Not able to found directory:", err)
			os.Exit(1)
		}
		s, err := store.Load(path)
		if err != nil {
			fmt.Println("Not able to load json:", err)
			os.Exit(1)
		}
		ui.Start(s)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
