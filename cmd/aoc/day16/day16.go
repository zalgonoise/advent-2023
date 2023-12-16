package day16

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d16 "github.com/zalgonoise/advent-2023/day16"
	p1 "github.com/zalgonoise/advent-2023/day16/part01"
	p2 "github.com/zalgonoise/advent-2023/day16/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-16", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the sum of energized tiles; starting on the top-left corner going to the right (east. Part 2 explores how many tiles are energized in the most optimal starting point. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d16.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 16),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		grid := p1.Parse(*input)
		result := p1.Sum(grid)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		grid := p2.Parse(*input)
		result := p2.Sum(grid)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
