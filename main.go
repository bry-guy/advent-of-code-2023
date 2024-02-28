package main

import (
	day1 "github.com/bry-guy/advent-of-code-2023/cmd/day-1"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc23",
	Short: "Advent of Code 2023 CLI",
}

func main() {
	rootCmd.Execute()
}

func init() {
	d1 := &cobra.Command{
		Use:   "day-1",
		Short: "Run day-1 challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day1.Run()
		},
	}

	// d2 := &cobra.Command{
	//     Use:   "day-2",
	//     Short: "Run day-2 challenge",
	//     Run: func(cmd *cobra.Command, args []string) {
	//         day2.Run()
	//     },
	// }

	rootCmd

	// rootCmd.AddCommand(d1, d2)
}
