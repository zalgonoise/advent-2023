package day08

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d8 "github.com/zalgonoise/advent-2023/day08"
	p1 "github.com/zalgonoise/advent-2023/day08/part01"
	p2 "github.com/zalgonoise/advent-2023/day08/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-08", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the number of steps to reach 'ZZZ'. An empty string uses the generated input")
	start := fs.String("start", "", "the starting node's value. The default is 'AAA' for part 1 or 'A' for part 2")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d8.Input
	}

	if *start == "" {
		switch *part {
		case 2:
			*start = "A"
		default:
			*start = "AAA"
		}
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 8),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		moves, nodes := p1.Parse(*input)
		result, err := p1.Find(*start, nodes, moves)
		if err != nil {
			return -1, err
		}

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		moves, nodes := p2.Parse(*input)
		result, err := p2.Find(*start, nodes, moves)
		if err != nil {
			return -1, err
		}

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
