package day13

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d13 "github.com/zalgonoise/advent-2023/day13"
	p1 "github.com/zalgonoise/advent-2023/day13/part01"
	p2 "github.com/zalgonoise/advent-2023/day13/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-13", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the sum of the mirroring points of the input, where as part 02 considers a single smudge in the calculation. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d13.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 13),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		fields := p1.Parse(*input)
		result := p1.Sum(fields)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		fields := p2.Parse(*input)
		result := p2.Sum(fields)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
