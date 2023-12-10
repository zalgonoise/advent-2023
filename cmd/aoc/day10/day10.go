package day10

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d10 "github.com/zalgonoise/advent-2023/day10"
	p1 "github.com/zalgonoise/advent-2023/day10/part01"
	p2 "github.com/zalgonoise/advent-2023/day10/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-10", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to solve the puzzles in the pipes maze. Part 1 calculates the largest ring's halfway distance while part 2 calculates the enclosed tiles in the maze. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d10.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 10),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		graph := p1.Parse(*input)
		result := p1.HalfwayDistance(graph)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		graph := p2.Parse(*input)
		result := p2.EnclosedTiles(graph)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
