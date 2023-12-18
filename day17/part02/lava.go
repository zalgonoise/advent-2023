package part02

import (
	"bufio"
	"strings"

	"github.com/zalgonoise/advent-2023/day17/grid"
)

const minAlloc = 64

func Parse(input string) grid.Graph[int] {
	lines := make([]string, 0, minAlloc)
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		lines = append(lines, line)
	}

	valueSet := make([][]int, 0, len(input))
	for i := range lines {
		values := make([]int, 0, len(lines[i]))

		for idx := range lines[i] {
			values = append(values, int(lines[i][idx]-'0'))
		}

		valueSet = append(valueSet, values)
	}

	m := grid.New(valueSet, grid.WithQuadrant(grid.Q3))

	return grid.Graph[int]{
		Head: grid.Coord{},
		Tail: grid.Coord{Y: m.MaxY, X: m.MaxX},
		Map:  m,
	}
}

func AStar(g grid.Graph[int]) int {
	return grid.AStar(g, 4, 10)
}
