package day21

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d19 "github.com/zalgonoise/advent-2023/day21"
	p1 "github.com/zalgonoise/advent-2023/day21/part01"
	p2 "github.com/zalgonoise/advent-2023/day21/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-21", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to count the steps when travelling the garden. An empty string uses the generated input")
	iter := fs.Int("iter", -1, "the number of steps to take. Defaults to 64 for part 1 and 26501365 for part 2.")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d19.Input
	}

	if *iter < 0 {
		switch *part {
		case 1:
			*iter = 64
		case 2:
			*iter = 26501365
		}
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 21),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		graph := p1.Parse(*input, *iter)
		result := p1.Count(graph)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		graph := p2.Parse(*input, *iter)
		result := p2.Count(graph)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
