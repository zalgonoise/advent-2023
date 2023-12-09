package part01

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

const (
	minAlloc  = 64
	tinyAlloc = 8
)

func Sum(valueSets [][]int) int {
	var n int

	for i := range valueSets {
		n += nextValue(valueSets[i])
	}

	return n
}

func nextValue(input []int) int {
	tree := make([][]int, 0, tinyAlloc)
	next := input
	var isZero bool

	for {
		next, isZero = diff(next)
		if isZero {
			break
		}

		tree = append(tree, next)
	}

	switch {
	case len(tree) == 1:
		return 0
	case len(tree) == 2 && allZeros(tree[1]):
		return input[len(input)-1] + tree[0][len(tree[0])-1]
	case !allZeros(tree[len(tree)-1]):
		return 0
	default:
		// set the last entry to match the streak, on the next-to-last index
		tree[len(tree)-2] = append(tree[len(tree)-2], tree[len(tree)-2][len(tree[len(tree)-2])-1])
	}

	for i := len(tree) - 3; i >= 0; i-- {
		tree[i] = append(tree[i], tree[i][len(tree[i])-1]+tree[i+1][len(tree[i+1])-1])
	}

	return input[len(input)-1] + tree[0][len(tree[0])-1]
}

func diff(input []int) (next []int, isZero bool) {
	if allZeros(input) {
		return input, true
	}

	out := make([]int, 0, len(input)) // leave the extra allocated space for the prediction

	for i := 1; i < len(input); i++ {
		out = append(out, input[i]-input[i-1])
	}

	return out, false
}

func allZeros(input []int) bool {
	var zeros int

	for i := range input {
		if input[i] == 0 {
			zeros++
		}
	}

	return zeros == len(input)
}

func Parse(input string) ([][]int, error) {
	if input == "" {
		return nil, nil
	}

	values := make([][]int, 0, minAlloc)
	errs := make([]error, 0, minAlloc)
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		lineValues := make([]int, 0, len(fields))

		for i := range fields {
			value, err := strconv.Atoi(fields[i])
			if err != nil {
				errs = append(errs, err)

				continue
			}

			lineValues = append(lineValues, value)
		}

		if len(lineValues) > 0 {
			values = append(values, lineValues)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return values, nil
}
