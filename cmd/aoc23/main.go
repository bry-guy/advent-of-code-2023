package main

import (
	"os"

	"github.com/bry-guy/advent-of-code-2023/pkg/day-1"
	"github.com/bry-guy/advent-of-code-2023/pkg/day-2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc23",
	Short: "aoc23 is advent of code 2023",
	Long:  `aoc23 is a CLI tool for running Advent of Code 2023 puzzle implementations.`,
}

func main() {
	rootCmd.PersistentFlags().Bool("debug", true, "enable debug logging")

	rootCmd.AddCommand(day1.NewCmd())
	rootCmd.AddCommand(day2.NewCmd())

	if err := rootCmd.Execute(); err != nil {
		// fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
