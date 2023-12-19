package day18

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d18 "github.com/zalgonoise/advent-2023/day18"
	p1 "github.com/zalgonoise/advent-2023/day18/part01"
	p2 "github.com/zalgonoise/advent-2023/day18/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-18", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the area of the lava pool. Part 1 calculates it according to the readable instructions in the input, while part 2 decodes the instructions from the hexadecimal number in the input. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d18.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 18),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		vectors, err := p1.Parse(*input)
		if err != nil {
			return 1, err
		}

		result := p1.Area(vectors)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		vectors, err := p2.Parse(*input)
		if err != nil {
			return 1, err
		}

		result := p2.Area(vectors)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
