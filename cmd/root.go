package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Advent of Go 2023 runner",
	Long:  `Runs an Advent of Code 2023 day's solution.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: go run . <day>")
		fmt.Println("Example: go run . 1")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
