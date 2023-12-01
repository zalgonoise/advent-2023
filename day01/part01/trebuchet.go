package part01

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
		if input[i] >= '0' && input[i] <= '9' {
			n = int(input[i]-'0') * 10

			break
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] >= '0' && input[i] <= '9' {
			n += int(input[i] - '0')

			break
		}
	}

	return n
}
