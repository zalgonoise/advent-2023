package part01

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	timePrefix = "Time:"
	distPrefix = "Distance:"
)

var (
	errInvalidPrefix = errors.New("invalid prefix")
	errLenMismatch   = errors.New("length mismatch")
)

type Race struct {
	total  int
	record int
}

func Parse(input string) ([]Race, error) {
	if input == "" {
		return nil, nil
	}

	var (
		times   []int
		records []int
		err     error
		errs    = make([]error, 0, 2)
		scanner = bufio.NewScanner(bytes.NewReader([]byte(input)))
	)

	for scanner.Scan() {
		line := string(scanner.Bytes())

		switch {
		case strings.HasPrefix(line, timePrefix):
			if times, err = getValues(line, timePrefix); err != nil {
				errs = append(errs, err)
			}

		case strings.HasPrefix(line, distPrefix):
			if records, err = getValues(line, distPrefix); err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(times) != len(records) {
		return nil, fmt.Errorf("%w: times: %d != records: %d", errLenMismatch, len(times), len(records))
	}

	races := make([]Race, 0, len(times))

	for i := range times {
		races = append(races, Race{total: times[i], record: records[i]})
	}

	return races, nil
}

func getValues(line, prefix string) ([]int, error) {
	split := strings.Split(line, prefix)
	if len(split) < 1 {
		return nil, fmt.Errorf("%w: no %q in %q", errInvalidPrefix, prefix, line)
	}

	values := strings.Fields(split[1])
	output := make([]int, 0, len(values))
	errs := make([]error, 0, len(values))

	for i := range values {
		value, err := strconv.Atoi(values[i])
		if err != nil {
			errs = append(errs, err)

			continue
		}

		output = append(output, value)
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return output, nil
}

func Sum(races ...Race) int {
	if len(races) == 0 {
		return 0
	}

	quantities := make([]int, 0, len(races))

	for i := range races {
		if total := Plan(races[i]); len(total) > 0 {
			quantities = append(quantities, len(total))
		}
	}

	if len(quantities) == 0 {
		return 0
	}

	var n = quantities[0]

	for i := 1; i < len(quantities); i++ {
		n *= quantities[i]
	}

	return n
}

func Plan(race Race) []int {
	attempts := make([]int, 0, race.total)

	for i := 1; i < race.total; i++ {
		if distance := boostDistance(i, race.total); distance > race.record {
			attempts = append(attempts, i)
		}
	}

	return attempts
}

func boostDistance(hold, limit int) int {
	if hold <= 0 || limit <= 0 || hold >= limit {
		return 0
	}

	return (limit - hold) * hold
}
