package cmd

import (
	"advent-of-go/day1"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day1Cmd)
	day1Cmd.Flags().StringP("file", "f", "./day1/input.txt", "Input file")
}

var day1Cmd = &cobra.Command{
	Use:   "1 --file <inputfile>",
	Short: "Run day1 solution",
	Long:  `Optionally specify input file. Default is input.txt.`,
	Run: func(cmd *cobra.Command, args []string) {
		day1.Execute(cmd.Flags().Lookup("file").Value.String())
	},
}
