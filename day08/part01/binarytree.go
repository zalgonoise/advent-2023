package part01

import (
	"bufio"
	"errors"
	"strings"
)

const (
	left  = 0
	right = 1

	minAlloc = 64
)

var (
	errNodeNotFound = errors.New("node not found")
	errStuckInLoop  = errors.New("got stuck in a loop")
)

func Find(start string, nodes map[string][2]string, moves []int) (int, error) {
	n, err := find(start, 0, nodes, moves)
	if err != nil {
		return -1, err
	}

	return n, nil
}

func find(start string, n int, nodes map[string][2]string, moves []int) (int, error) {
	key, err := findInSet(start, 0, nodes, moves)
	if err != nil {
		return -1, err
	}

	n += len(moves)

	if key != "ZZZ" {
		return find(key, n, nodes, moves)
	}

	return n, nil
}

func findInSet(key string, idx int, nodes map[string][2]string, moves []int) (string, error) {
	if idx >= len(moves) {
		return key, nil
	}

	v, ok := nodes[key]
	if !ok {
		return "", errNodeNotFound
	}

	if v[0] == key && v[1] == key {
		return "", errStuckInLoop
	}

	idx++
	return findInSet(v[moves[idx-1]], idx, nodes, moves)
}

func Parse(input string) ([]int, map[string][2]string) {
	if input == "" {
		return nil, nil
	}

	var (
		moves []int
		nodes = make(map[string][2]string, minAlloc)
	)

	scanner := bufio.NewScanner(strings.NewReader(input))

	if scanner.Scan() {
		moves = parseMoves(scanner.Text())
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		key, l, r := parseLine(line)
		nodes[key] = [2]string{l, r}
	}

	return moves, nodes
}

func parseMoves(line string) []int {
	moves := make([]int, 0, len(line))

	for i := range line {
		switch line[i] {
		case 'L':
			moves = append(moves, left)
		case 'R':
			moves = append(moves, right)
		}
	}

	return moves
}

func parseLine(line string) (key, left, right string) {
	fields := strings.Fields(line)

	if len(fields) != 4 {
		return "", "", ""
	}

	return getField(fields[0]), getField(fields[2]), getField(fields[3])
}

func getField(field string) string {
	out := make([]byte, 0, len(field))

	for i := range field {
		if field[i] < 'A' || field[i] > 'Z' {
			continue
		}

		out = append(out, field[i])
	}

	return string(out)
}
