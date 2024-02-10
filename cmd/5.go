package cmd

import (
	"advent-of-go/day5"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day5Cmd)
	day5Cmd.Flags().StringP("file", "f", "./day5/input.txt", "Input file")
}

var day5Cmd = &cobra.Command{
	Use:   "5 --file <inputfile>",
	Short: "Run day5 solution",
	Long:  `Optionally specify input file. Default is input.txt.`,
	Run: func(cmd *cobra.Command, args []string) {
		day5.Execute(cmd.Flags().Lookup("file").Value.String())
	},
}
