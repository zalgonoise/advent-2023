package part01

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var errInvalidNumFields = errors.New("invalid number of fields")

type Set struct {
	items  string
	combos []int
}

func Sum(sets []Set) int {
	var (
		sum   int
		cache = map[string]int{}
	)

	for i := range sets {
		sum += explore(sets[i].items, sets[i].combos, -1, cache)
	}

	return sum
}

func explore(row string, groups []int, cur int, cache map[string]int) int {
	if sum, ok := cache[key(row, groups, cur)]; ok {
		return sum
	}

	switch {
	case row == "" && len(groups) == 0 && cur <= 0:
		return 1
	case row == "":
		return 0
	}

	var n int

	switch {
	case row[0] == '#':
		switch {
		case cur == 0, cur == -1 && len(groups) == 0:
			return 0
		case cur == -1:
			cur = groups[0]
			groups = groups[1:]
		}

		n = explore(row[1:], groups, cur-1, cache)
	case row[0] == '.' && cur <= 0:
		n = explore(row[1:], groups, -1, cache)
	case row[0] == '?':
		n = explore("#"+row[1:], groups, cur, cache) + explore("."+row[1:], groups, cur, cache)
	}

	cache[key(row, groups, cur)] = n

	return n
}

func key(row string, groups []int, currentGroup int) string {
	return fmt.Sprintf("%s::%v::%d", row, groups, currentGroup)
}

func Parse(input string, factor int) ([]Set, error) {
	if input == "" {
		return nil, nil
	}

	lines := strings.Split(input, "\n")
	sets := make([]Set, 0, len(lines))

	for i := range lines {
		if lines[i] == "" {
			continue
		}

		set, err := extract(lines[i], factor)
		if err != nil {
			return nil, err
		}

		sets = append(sets, set)
	}

	return sets, nil
}

func makeCopy[T any, S ~[]T](n int, slice S) S {
	s := make(S, len(slice)*n)
	idx := 0

	for i := 0; i < n; i++ {
		idx += copy(s[idx:], slice)
	}

	return s
}

func extract(line string, factor int) (Set, error) {
	fields := strings.Fields(line)

	if len(fields) != 2 {
		return Set{}, errInvalidNumFields
	}

	combosValues := strings.Split(fields[1], ",")
	combos := make([]int, 0, len(combosValues))
	errs := make([]error, 0, len(combosValues))

	for i := range combosValues {
		combo, err := strconv.Atoi(combosValues[i])
		if err != nil {
			errs = append(errs, err)

			continue
		}

		combos = append(combos, combo)
	}

	combos = makeCopy(factor, combos)
	sb := &strings.Builder{}

	for i := 0; i < factor; i++ {
		sb.WriteString(fields[0])

		if i < factor-1 {
			sb.WriteByte('?')
		}
	}

	if len(errs) > 0 {
		return Set{}, errors.Join(errs...)
	}

	return Set{
		items:  sb.String(),
		combos: combos,
	}, nil
}
