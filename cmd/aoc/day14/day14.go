package day14

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d14 "github.com/zalgonoise/advent-2023/day14"
	p1 "github.com/zalgonoise/advent-2023/day14/part01"
	p2 "github.com/zalgonoise/advent-2023/day14/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-14", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the sum of weight on the north parabolic dish structure, whereas part 2 predicts how that pressure will be in X iterations where the rocks are rotated on all four axis. An empty string uses the generated input")
	iter := fs.Int("iter", 0, "the number of iterations for each tilt action, applicable for part 2.")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d14.Input
	}

	if *iter == 0 {
		*iter = 1000000000
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 14),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		grid := p1.Parse(*input)
		grid.Tilt()

		result := grid.Sum()
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		grid := p2.Parse(*input)
		result := grid.Sum(*iter)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
