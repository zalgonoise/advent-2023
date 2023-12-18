package day17

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"

	d17 "github.com/zalgonoise/advent-2023/day17"
	p1 "github.com/zalgonoise/advent-2023/day17/part01"
	p2 "github.com/zalgonoise/advent-2023/day17/part02"
)

var (
	errInvalidPart     = errors.New("invalid part")
	errUnsupportedPart = errors.New("unsupported part")
)

func Exec(ctx context.Context, logger *slog.Logger, args []string) (int, error) {
	fs := flag.NewFlagSet("day-17", flag.ExitOnError)

	part := fs.Int("part", 1, "either part 1 or 2")
	input := fs.String("input", "", "the input string to calculate the total heat loss going from the top-left to the bottom-right corner of the grid. Part 1 considers a minimum of zero steps and maximum of 3. Part 2 considers a minimum of 4 steps and maximum of 10. An empty string uses the generated input")

	if err := fs.Parse(args); err != nil {
		return 1, err
	}

	if *part < 1 || *part > 2 {
		return 1, fmt.Errorf("%w: %d", errInvalidPart, *part)
	}

	if *input == "" {
		*input = d17.Input
	}

	attr := slog.Group("challenge",
		slog.String("name", "Advent of Code 2023"),
		slog.Int("day", 17),
		slog.Int("part", *part),
	)

	switch *part {
	case 1:
		graph := p1.Parse(*input)
		result := p1.AStar(graph)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	case 2:
		graph := p2.Parse(*input)
		result := p2.AStar(graph)

		logger.InfoContext(ctx, "execution completed", slog.Int("result", result), attr)
	default:
		return 1, fmt.Errorf("%w: %d", errUnsupportedPart, *part)
	}

	return 0, nil
}
