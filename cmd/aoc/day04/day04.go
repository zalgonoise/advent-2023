package day04

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d4 "github.com/zalgonoise/advent-2023/day04"
	p1 "github.com/zalgonoise/advent-2023/day04/part01"
	p2 "github.com/zalgonoise/advent-2023/day04/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-04", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the total points in the scratchcards. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d4.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 4),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		result := p1.Sum(p1.NewSet(*input)...)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		result := p2.Sum(p2.NewSet(*input)...)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
