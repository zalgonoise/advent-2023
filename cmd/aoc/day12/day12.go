package day12

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d12 "github.com/zalgonoise/advent-2023/day12"
	p1 "github.com/zalgonoise/advent-2023/day12/part01"
	p2 "github.com/zalgonoise/advent-2023/day12/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-12", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the sum of the potential combinations for the broken and running sprints. An empty string uses the generated input")
	factor := fs.Int("factor", 0, "the folding factor for the springs' input. Defaults are 1 for part 1 and 5 for part 2.")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d12.Input
	}

	if *factor == 0 {
		switch *part {
		case 2:
			*factor = 5
		default:
			*factor = 1
		}
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 12),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		sets, err := p1.Parse(*input, *factor)
		if err != nil {
			return 1, err
		}

		result := p1.Sum(sets)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		sets, err := p2.Parse(*input, *factor)
		if err != nil {
			return 1, err
		}

		result := p2.Sum(sets)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
