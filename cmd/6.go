package cmd

import (
	"advent-of-go/day6"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day6Cmd)
	day6Cmd.Flags().StringP("file", "f", "./day6/input.txt", "Input file")
}

var day6Cmd = &cobra.Command{
	Use:   "6 --file <inputfile>",
	Short: "Run day6 solution",
	Long:  `Optionally specify input file. Default is input.txt.`,
	Run: func(cmd *cobra.Command, args []string) {
		day6.Execute(cmd.Flags().Lookup("file").Value.String())
	},
}
