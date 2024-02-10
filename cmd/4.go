package cmd

import (
	"advent-of-go/day4"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day4Cmd)
	day4Cmd.Flags().StringP("file", "f", "./day4/input.txt", "Input file")
}

var day4Cmd = &cobra.Command{
	Use:   "4 --file <inputfile>",
	Short: "Run day4 solution",
	Long:  `Optionally specify input file. Default is input.txt.`,
	Run: func(cmd *cobra.Command, args []string) {
		day4.Execute(cmd.Flags().Lookup("file").Value.String())
	},
}
