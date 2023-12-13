package part01

import (
	"bufio"
	"strings"
)

const minAlloc = 16

type Field struct {
	rows []string
}

func (f Field) Sum() int {
	// find mirrors in Y axis
	n := find(f.rows) * 100

	// find mirrors in X axis
	n += find(rotate(f.rows))

	return n
}

func Sum(fields []Field) int {
	var n int

	for i := range fields {
		n += fields[i].Sum()
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

func find(rows []string) int {
	caught := make([]int, 0, len(rows))

	for i := 1; i < len(rows); i++ {
		if rows[i] == rows[i-1] {
			caught = append(caught, i-1)
		}
	}

	for idx := range caught {
		isMatch := true
		for i, j := caught[idx], caught[idx]+1; i >= 0; i, j = i-1, j+1 {
			if j >= len(rows) {
				break
			}

			if rows[i] != rows[j] {
				isMatch = false

				break
			}
		}

		if isMatch {
			return caught[idx] + 1
		}
	}

	return 0
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
