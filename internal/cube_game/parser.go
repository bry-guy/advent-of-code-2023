package cube_game

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

/*
A game is possible if, for every set of cubes shown in that game,
the number of cubes of each color does not exceed the total number of cubes of that color available in the bag.

Example:
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
*/

type gameResult struct {
	blue  int
	red   int
	green int
}

func valid(game gameResult) bool {
	return game.blue <= 14 && game.red <= 12 && game.green <= 13
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func Parse(gameSet string) (int, error) {
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, &opts)
	log := slog.New(handler)

	// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	gameId, gamesStr, ok := strings.Cut(gameSet, ":")
	if !ok {
		return 0, fmt.Errorf("unable to Parse: missing semicolon: str: %s", gameSet)
	}

	// Game 1
	_, idStr, ok := strings.Cut(gameId, " ")
	if !ok {
		return 0, fmt.Errorf("unable to Parse: cannot find id: str: %s", gameId)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("unable to Parse: cannot convert id: str: %s", idStr)
	}

	// [ "3 blue, 4 red", "1 red, 2 green" ]
	games := strings.Split(gamesStr, ";")
	if len(games) < 1 {
		return 0, fmt.Errorf("unable to Parse: no games found: str: %s", gamesStr)
	}

	var gameResult gameResult
	for _, game := range games {
		gameResult, err = parse(game)
		if err != nil {
			return 0, fmt.Errorf("unable to Parse: %w", err)
		}

		if !valid(gameResult) {
			log.Debug("invalid", "gameResult", gameResult)
			return 0, nil
		}
	}

	log.Debug("valid", "gameResult", gameResult)
	return id, nil

}

func parse(game string) (gameResult, error) {
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, &opts)
	log := slog.New(handler)

	var gameResult gameResult

	// "3 blue"
	countColors := strings.Split(game, ",")

	// ["3 blue", "4 red"]
	// actual: ["3 blue 4 red"]
	for _, countColor := range countColors {
		cc := strings.TrimSpace(countColor)
		log.Debug("parse count color", "cc", cc)
		countStr, color, ok := strings.Cut(cc, " ")
		if !ok {
			return gameResult, fmt.Errorf("unable to find count and color: str: %s", countColor)
		}

		count, err := strconv.Atoi(countStr)
		if err != nil {
			return gameResult, fmt.Errorf("invalid count: str: %s", countStr)
		}

		switch color {
		case "blue":
			gameResult.blue += count
		case "red":
			gameResult.red += count
		case "green":
			gameResult.green += count
		default:
			return gameResult, fmt.Errorf("invalid color found: str: %s", color)
		}
	}

	return gameResult, nil
}
