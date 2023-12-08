package part02

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

var errNodeNotFound = errors.New("node not found")

func Find(start string, nodes map[string][2]string, moves []int) (int, error) {
	startNodes := findStartNodes(start, nodes)

	result := 1
	for i := range startNodes {
		n, err := find(startNodes[i], 0, nodes, moves)
		if err != nil {
			return -1, err
		}

		result = leastMultiple(result, n)
	}

	return result, nil
}

func greatestDivisor(a, b int) int {
	if b == 0 {
		return a
	}

	return greatestDivisor(b, a%b)
}

func leastMultiple(a int, b int) int {
	return (a / greatestDivisor(a, b)) * b
}

func findStartNodes(target string, nodes map[string][2]string) []string {
	keys := make([]string, 0, len(nodes))

	for k := range nodes {
		if strings.HasSuffix(k, target) {
			keys = append(keys, k)
		}
	}

	return keys
}

func find(start string, n int, nodes map[string][2]string, moves []int) (int, error) {
	key, err := findInSet(start, 0, nodes, moves)
	if err != nil {
		return -1, err
	}

	n += len(moves)

	if !strings.HasSuffix(key, "Z") {
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
		if (field[i] < 'A' || field[i] > 'Z') && (field[i] < '0' || field[i] > '9') {
			continue
		}

		out = append(out, field[i])
	}

	return string(out)
}
