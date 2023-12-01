package trebuchet

import (
	"strings"
)

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
