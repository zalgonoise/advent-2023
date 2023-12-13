package part02

import (
	"bufio"
	"strings"
)

const minAlloc = 16

type Field struct {
	rows []string
}

func Parse(input string) []Field {
	if input == "" {
		return nil
	}

	blocks := strings.Split(input, "\n\n")

	fields := make([]Field, 0, len(blocks))

	for i := range blocks {
		rows := make([]string, 0, minAlloc)

		scanner := bufio.NewScanner(strings.NewReader(blocks[i]))

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			rows = append(rows, line)
		}

		fields = append(fields, Field{rows: rows})
	}

	return fields
}

func find(lines []string) int {
	for i := 1; i < len(lines); i++ {
		isMatch := true
		numSmudges := 0
		lim := min(i, len(lines)-i)

		for idx := 0; isMatch && idx < lim; idx++ {
			if lines[i-idx-1] == lines[i+idx] {
				continue
			}

			// short-circuit out
			if numSmudges > 0 {
				isMatch = false

				break
			}

			matchingBytes := len(lines[i+idx])
			for ii := range lines[i-idx-1] {
				if lines[i-idx-1][ii] == lines[i+idx][ii] {
					matchingBytes--
				}
			}

			switch matchingBytes {
			case 0:
			case 1:
				numSmudges++
			default:
				isMatch = false
			}
		}

		if isMatch && numSmudges == 1 {
			return i
		}
	}

	return 0
}

func Sum(fields []Field) int {
	var n int

	for i := range fields {
		n += find(rotate(fields[i].rows))
		n += find(fields[i].rows) * 100
	}

	return n
}

func rotate(input []string) []string {
	rotated := make([]string, 0, len(input[0]))

	for i := range input[0] {
		temp := make([]byte, len(input))

		for idx := range input {
			temp[idx] = input[idx][i]
		}

		rotated = append(rotated, string(temp))
	}

	return rotated
}
