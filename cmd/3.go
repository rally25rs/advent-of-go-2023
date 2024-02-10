package cmd

import (
	"advent-of-go/day3"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day3Cmd)
	day3Cmd.Flags().StringP("file", "f", "./day3/input.txt", "Input file")
}

var day3Cmd = &cobra.Command{
	Use:   "3 --file <inputfile>",
	Short: "Run day3 solution",
	Long:  `Optionally specify input file. Default is input.txt.`,
	Run: func(cmd *cobra.Command, args []string) {
		day3.Execute(cmd.Flags().Lookup("file").Value.String())
	},
}
