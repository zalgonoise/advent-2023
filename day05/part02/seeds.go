package part02

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	bigAlloc  = 1_000_000
	minAlloc  = 32
	numFields = 3
)

type Plot []Mapping

type Mapping struct {
	to   int
	from int
	span int
}

func Parse(input string) ([]int, []Plot, error) {
	if input == "" {
		return nil, nil, nil
	}

	var (
		seeds []int
		plots = make([]Plot, 0, 7) // expects 7 mappings from example and input
		lines = strings.Split(input, "\n")
		errs  = make([]error, 0, minAlloc)
	)

	for i := 0; i < len(lines); i++ {
		switch {
		case strings.HasPrefix(lines[i], "seeds: "):
			seedValues := strings.Fields(strings.Split(lines[i], "seeds: ")[1])
			seeds = make([]int, 0, bigAlloc)

			for idx := 0; idx < len(seedValues); idx += 2 {
				if idx >= len(seedValues) || idx+1 >= len(seedValues) {
					break
				}

				from, err := strconv.Atoi(seedValues[idx])
				if err != nil {
					errs = append(errs, err)

					continue
				}

				span, err := strconv.Atoi(seedValues[idx+1])
				if err != nil {
					errs = append(errs, err)

					continue
				}

				for ; span > 0; span-- {
					seeds = append(seeds, from)
					from++
				}
			}

		case strings.HasPrefix(lines[i], "seed-to-soil map:"),
			strings.HasPrefix(lines[i], "soil-to-fertilizer map:"),
			strings.HasPrefix(lines[i], "fertilizer-to-water map:"),
			strings.HasPrefix(lines[i], "water-to-light map:"),
			strings.HasPrefix(lines[i], "light-to-temperature map:"),
			strings.HasPrefix(lines[i], "temperature-to-humidity map:"),
			strings.HasPrefix(lines[i], "humidity-to-location map:"):
			plot, n, err := newPlot(lines[i+1:])

			i += n
			if err != nil {
				errs = append(errs, err)
			}

			plots = append(plots, plot)
		}
	}

	return seeds, plots, errors.Join(errs...)
}

func newPlot(lines []string) (Plot, int, error) {
	var (
		i    = 0
		plot = make(Plot, 0, minAlloc)
		errs = make([]error, 0, len(lines))
	)

	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}

		values := strings.Fields(lines[i])
		if len(values) != numFields {
			errs = append(errs, fmt.Errorf("mapping with multiple fields: %d", len(values)))

			continue
		}

		fields := make([]int, 0, numFields)

		for ii := range values {
			value, err := strconv.Atoi(values[ii])
			if err != nil {
				errs = append(errs, err)

				break
			}

			fields = append(fields, value)
		}

		if len(fields) != numFields {
			errs = append(errs, fmt.Errorf("failed to decode all fields: decoded %d", len(fields)))

			continue
		}

		plot = append(plot, Mapping{
			to:   fields[0],
			from: fields[1],
			span: fields[2],
		})
	}

	return plot, i + 1, errors.Join(errs...)
}

func (m Mapping) Plant(seed int) int {
	pos := seed - m.from

	if pos < 0 || pos > m.span {
		return -1
	}

	return m.to + pos
}

func (p Plot) Plant(seed int) int {
	if len(p) == 0 {
		return seed
	}

	n := -1

	for i := range p {
		pos := p[i].Plant(seed)
		if pos == -1 {
			continue
		}

		n = pos

		break
	}

	if n < 0 {
		return seed
	}

	return n
}

func Plant(seed int, plots ...Plot) int {
	if len(plots) == 0 {
		return seed
	}

	for i := range plots {
		seed = plots[i].Plant(seed)
	}

	return seed
}

func Lowest(seeds []int, plots ...Plot) int {
	if len(seeds) == 0 {
		return -1
	}

	if len(plots) == 0 {
		slices.Sort(seeds)

		return seeds[0]
	}

	positions := make([]int, 0, len(seeds))

	for i := range seeds {
		positions = append(positions, Plant(seeds[i], plots...))
	}

	return slices.Min(positions)
}
