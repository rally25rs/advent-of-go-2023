package cmd

import (
	"advent-of-go/day2"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day2Cmd)
	day2Cmd.Flags().StringP("file", "f", "./day2/input.txt", "Input file")
}

var day2Cmd = &cobra.Command{
	Use:   "2 --file <inputfile>",
	Short: "Run day2 solution",
	Long:  `Optionally specify input file. Default is input.txt.`,
	Run: func(cmd *cobra.Command, args []string) {
		day2.Execute(cmd.Flags().Lookup("file").Value.String())
	},
}
