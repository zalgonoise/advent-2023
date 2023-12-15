package part01

import "strings"

func Parse(input string) []string {
	if input == "" {
		return nil
	}

	input = strings.ReplaceAll(input, "\n", "")
	raw := strings.Split(input, ",")

	fields := make([]string, 0, len(raw))

	for i := range raw {
		if raw[i] == "" {
			continue
		}

		fields = append(fields, raw[i])
	}

	return fields
}

func HashSum(input ...string) int {
	var n int

	for i := range input {
		n += Hash(input[i])
	}

	return n
}

func Hash(input string) int {
	var n int

	for i := range input {
		n = hash(n + int(input[i]))
	}

	return n
}

func hash(value int) int {
	value *= 17
	return value % 256
}
