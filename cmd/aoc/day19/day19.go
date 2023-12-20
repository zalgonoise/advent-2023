package day19

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d19 "github.com/zalgonoise/advent-2023/day19"
	p1 "github.com/zalgonoise/advent-2023/day19/part01"
	p2 "github.com/zalgonoise/advent-2023/day19/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-19", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to organize the parts for. Part 2 discards the provided ranges and spans from a minimum and maximum value. An empty string uses the generated input")
	from := fs.Int("from", -1, "the value to span from when scanning a range, for the 2nd part of the challenge. A negative value defaults to 1.")
	to := fs.Int("to", -1, "the value to span to when scanning a range, for the 2nd part of the challenge. A negative value defaults to 4000.")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d19.Input
	}

	if *from < 0 {
		*from = 1
	}

	if *to < 0 || *to < *from {
		*to = 4000
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 19),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		pipeline, moves, err := p1.Parse(*input)
		if err != nil {
			return 1, err
		}

		outcome := pipeline.Travel(moves)

		result := p1.Sum(outcome)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		pipeline, err := p2.Parse(*input)
		if err != nil {
			return 1, err
		}

		result := pipeline.Travel(p2.Range[p2.Part]{
			Min: p2.Part{*from, *from, *from, *from},
			Max: p2.Part{*to, *to, *to, *to},
		})

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
