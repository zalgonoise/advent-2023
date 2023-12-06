package part02

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

func Parse(input string) (Race, error) {
	if input == "" {
		return Race{}, nil
	}

	var (
		err     error
		race    = Race{}
		scanner = bufio.NewScanner(bytes.NewReader([]byte(input)))
	)

	for scanner.Scan() {
		line := string(scanner.Bytes())

		switch {
		case strings.HasPrefix(line, timePrefix):
			if race.total, err = getValues(line, timePrefix); err != nil {
				return Race{}, err
			}

		case strings.HasPrefix(line, distPrefix):
			if race.record, err = getValues(line, distPrefix); err != nil {
				return Race{}, err
			}
		}
	}

	return race, nil
}

func getValues(line, prefix string) (int, error) {
	split := strings.Split(line, prefix)
	if len(split) < 1 {
		return -1, fmt.Errorf("%w: no %q in %q", errInvalidPrefix, prefix, line)
	}

	values := strings.Fields(split[1])
	value := strings.Join(values, "")

	v, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}

	return v, nil
}

func Plan(race Race) int {
	var n int

	for i := 1; i < race.total; i++ {
		if distance := boostDistance(i, race.total); distance > race.record {
			n++
		}
	}

	return n
}

func boostDistance(hold, limit int) int {
	if hold <= 0 || limit <= 0 || hold >= limit {
		return 0
	}

	return (limit - hold) * hold
}
