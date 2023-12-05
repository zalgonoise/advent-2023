package day05

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d5 "github.com/zalgonoise/advent-2023/day05"
	p1 "github.com/zalgonoise/advent-2023/day05/part01"
	p2 "github.com/zalgonoise/advent-2023/day05/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-05", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the smallest resulting seed plot number. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d5.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 5),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		seeds, plots, err := p1.Parse(*input)
		if err != nil {
			return -1, err
		}

		result := p1.Lowest(seeds, plots...)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		seeds, plots, err := p2.Parse(*input)
		if err != nil {
			return -1, err
		}

		result := p2.Lowest(seeds, plots...)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
