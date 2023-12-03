package day03

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d3 "github.com/zalgonoise/advent-2023/day03"
	p1 "github.com/zalgonoise/advent-2023/day03/part01"
	p2 "github.com/zalgonoise/advent-2023/day03/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-03", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the gear ratios. An empty string uses the generated input")
	distance := fs.Int("range", 1, "range when a symbol connects two numbers. default is 1.")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d3.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 3),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		matrix := p1.NewMatrix(*input)
		coords := matrix.NearSymbol(*distance)
		result, err := matrix.Sum(coords)
		if err != nil {
			return -1, err
		}

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		matrix := p2.NewMatrix(*input)
		coords := matrix.NearSymbol(*distance)
		result, err := matrix.Sum(coords)
		if err != nil {
			return -1, err
		}

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
