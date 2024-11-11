package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobag",
	Short: "GoBag is a collection of generally useful cli tools",
	Long:  "GoBag is a collection of general, yet useful, command-line tools written in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to GoBag! Use --help for usage.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Uh-oh! An error occurred while executing gobag: '%s'\n", err)
		os.Exit(1)
	}
}

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s from Git SHA %s)", version, date, commit)
}
