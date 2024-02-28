package day2

/*
Given several games,
each with multiple sets of red, green, and blue cubes revealed from a bag,
determine which games could occur with a bag containing exactly 12 red, 13 green, and 14 blue cubes.

A game is possible if,
for every set of cubes shown in that game,
the number of cubes of each color does not exceed the total number of cubes of that color available in the bag.

Calculate the sum of the IDs of all possible games.

Example:
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
*/

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/bry-guy/advent-of-code-2023/internal/cube_game"
	"github.com/spf13/cobra"
)

type options struct {
	filepath string
	async    bool
}

func NewCmd() *cobra.Command {
	var opts options

	cmd := &cobra.Command{
		Use:   "day2 -f/--filepath {path-to-input} -a/--async",
		Short: "Run day2",
		Long:  `Run day2`,
		Run: func(cmd *cobra.Command, args []string) {
			d2(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.filepath, "filepath", "f", "", "filepath to input (required)")
	err := cmd.MarkFlagRequired("filepath")
	if err != nil {
		slog.Error("unable to mark filepath required", "err", err)
		os.Exit(1)
	}

	cmd.Flags().BoolVarP(&opts.async, "async", "a", false, "enable async processing")

	return cmd
}

func d2(opts *options) {
	start := time.Now()

	ctx := context.Background()
	log := slog.Default()

	log.InfoContext(ctx, "Advent of Code 2023, Day 2!")

	buf, err := os.ReadFile(opts.filepath)
	if err != nil {
		fmt.Printf("error: unable to read file at %s\n", opts.filepath)
	}

	reader := bufio.NewReader(bytes.NewReader(buf))

	var games []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		l := strings.TrimSuffix(line, "\n")

		games = append(games, l)
	}

	var sum, val int
	for _, game := range games {
		val, err = cube_game.Parse(game)
		if err != nil {
			log.ErrorContext(ctx, "unable to parse game", "game", game, "err", err)
			os.Exit(1)
		}
		sum += val
	}

	log.InfoContext(ctx, "parsing input complete")
	fmt.Println(sum)

	end := time.Now()
	duration := end.UnixMicro() - start.UnixMicro()

	fmt.Printf("duration: %v us\n", duration)
}
