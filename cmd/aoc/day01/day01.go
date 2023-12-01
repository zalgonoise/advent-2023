package day01

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d1 "github.com/zalgonoise/advent-2023/day01"
	p1 "github.com/zalgonoise/advent-2023/day01/part01"
	p2 "github.com/zalgonoise/advent-2023/day01/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-01", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the calibration total. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d1.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 1),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		result := p1.Trebuchet(*input)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		result := p2.Trebuchet(*input)
		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
