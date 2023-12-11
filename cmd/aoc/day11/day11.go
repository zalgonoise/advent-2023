package day11

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d11 "github.com/zalgonoise/advent-2023/day11"
	p1 "github.com/zalgonoise/advent-2023/day11/part01"
	p2 "github.com/zalgonoise/advent-2023/day11/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-11", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the sum of the manhattan distance between the galaxies, after the expansion. If unset, part 1 uses a factor of '2' while part 2 uses a factor of '1000000'. An empty string uses the generated input")
	factor := fs.Int("factor", 0, "the expansion factor when galaxies are spread apart")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d11.Input
	}

	if *factor == 0 {
		switch *part {
		case 2:
			*factor = 1_000_000
		default:
			*factor = 2
		}
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 11),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		grid := p1.Parse(*input)
		result := p1.Sum(grid, *factor)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		grid := p2.Parse(*input)
		result := p2.Sum(grid, *factor)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
