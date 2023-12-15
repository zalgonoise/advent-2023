package day15

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d15 "github.com/zalgonoise/advent-2023/day15"
	p1 "github.com/zalgonoise/advent-2023/day15/part01"
	p2 "github.com/zalgonoise/advent-2023/day15/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-15", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the sum of the hash in every ASCII character of the input, separated by commas; or to calculate the focal strength of the lenses if part 2 is selected. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d15.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 15),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		steps := p1.Parse(*input)
		result := p1.HashSum(steps...)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		ops := p2.Parse(*input)
		m := p2.Map(ops...)
		result := p2.Sum(m)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
