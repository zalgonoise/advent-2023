package day02

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d2 "github.com/zalgonoise/advent-2023/day02"
	p1 "github.com/zalgonoise/advent-2023/day02/part01"
	p2 "github.com/zalgonoise/advent-2023/day02/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-02", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the calibration total. An empty string uses the generated input")
	red := fs.Int("red", 0, "red value limit, applicable for part 1")
	green := fs.Int("green", 0, "green value limit, applicable for part 1")
	blue := fs.Int("blue", 0, "blue value limit, applicable for part 1")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d2.Input
	}

	if *red > 0 || *green > 0 || *blue > 0 {
		*part = 1
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 2),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		result, err := p1.CubeConundrum(*input, *red, *green, *blue)
		if err != nil {
			return 1, err
		}

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		result, err := p2.CubeConundrum(*input)
		if err != nil {
			return 1, err
		}

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
