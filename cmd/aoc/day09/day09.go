package day09

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d9 "github.com/zalgonoise/advent-2023/day09"
	p1 "github.com/zalgonoise/advent-2023/day09/part01"
	p2 "github.com/zalgonoise/advent-2023/day09/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-09", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the sum of the predictions in the input. Part 1 sums the future predictions while part 2 adds guessed past values. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d9.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 9),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		values, err := p1.Parse(*input)
		if err != nil {
			return -1, err
		}

		result := p1.Sum(values)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		values, err := p2.Parse(*input)
		if err != nil {
			return -1, err
		}

		result := p2.Sum(values)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
