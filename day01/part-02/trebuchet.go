package trebuchet

import (
	"strings"
)

var numbers = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

var numStrings = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func Trebuchet(input ...string) int {
	// short-circuit if empty
	if len(input) == 0 {
		return 0
	}

	var n int

	for i := range input {
		// split by newlines if any
		split := strings.Split(input[i], "\n")

		for idx := range split {
			n += trebuchet(split[idx])
		}
	}

	return n
}

func trebuchet(input string) int {
	var n int

	for i := range input {
		if value, ok := numOrString(input[i:]); ok {
			n = value * 10

			break
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		if value, ok := numOrString(input[i:]); ok {
			n += value

			break
		}
	}

	return n
}

func numOrString(input string) (int, bool) {
	switch {
	case input[0] >= '0' && input[0] <= '9':
		return int(input[0] - '0'), true
	case strings.HasPrefix(input, "one"):
		return 1, true
	case strings.HasPrefix(input, "two"):
		return 2, true
	case strings.HasPrefix(input, "three"):
		return 3, true
	case strings.HasPrefix(input, "four"):
		return 4, true
	case strings.HasPrefix(input, "five"):
		return 5, true
	case strings.HasPrefix(input, "six"):
		return 6, true
	case strings.HasPrefix(input, "seven"):
		return 7, true
	case strings.HasPrefix(input, "eight"):
		return 8, true
	case strings.HasPrefix(input, "nine"):
		return 9, true
	default:
		return 0, false
	}
}
