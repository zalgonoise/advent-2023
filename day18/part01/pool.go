package part01

import (
	"errors"
	"strconv"
	"strings"

	"github.com/zalgonoise/advent-2023/day18/grid"
)

var (
	errInvalidDirection = errors.New("invalid direction")
)

func Parse(input string) ([]grid.Vector, error) {
	if input == "" {
		return nil, nil
	}

	lines := strings.Split(input, "\n")
	vectors := make([]grid.Vector, 0, len(lines))

	for i := range lines {
		if lines[i] == "" {
			continue
		}

		move, err := extract(lines[i])
		if err != nil {
			return nil, err
		}

		vectors = append(vectors, move)
	}

	return vectors, nil
}

func extract(line string) (grid.Vector, error) {
	fields := strings.Fields(line)
	m := grid.Vector{}
	var err error

	switch fields[0][0] {
	case 'U':
		m.Dir = grid.South
	case 'D':
		m.Dir = grid.North
	case 'L':
		m.Dir = grid.West
	case 'R':
		m.Dir = grid.East
	default:
		return grid.Vector{}, errInvalidDirection
	}

	m.Len, err = strconv.Atoi(fields[1])
	if err != nil {
		return grid.Vector{}, err
	}

	return m, nil
}

func travel(start grid.Coord, vectors []grid.Vector) []grid.Coord {
	cur := start
	coords := make([]grid.Coord, 0, len(vectors)+1)

	for i := range vectors {
		coords = append(coords, cur)
		cur = grid.Add(cur, grid.Mul(vectors[i].Dir, vectors[i].Len))
	}

	coords = append(coords, cur)

	return coords
}

func shoelace(vertices []grid.Coord) int {
	var n int

	for i := range vertices {
		next := (i + 1) % len(vertices)

		n += vertices[i].X * vertices[next].Y
		n -= vertices[i].Y * vertices[next].X
	}

	return grid.Abs(n) / 2
}

func perimeter(vertices []grid.Coord) int {
	var n int

	for i := 0; i < len(vertices); i++ {
		next := (i + 1) % len(vertices)

		sub := grid.Sub(vertices[i], vertices[next])
		n += grid.Abs(sub.X) + grid.Abs(sub.Y)
	}

	return n
}

func Area(vectors []grid.Vector) int {
	coords := travel(grid.Coord{}, vectors)

	return shoelace(coords) + perimeter(coords)/2 + 1
}
