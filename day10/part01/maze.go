package part01

import (
	"bufio"
	"strings"

	"github.com/zalgonoise/advent-2023/day10/graph"
)

type Coord struct {
	Y int
	X int
}

func add(a, b Coord) Coord {
	return Coord{
		Y: a.Y + b.Y,
		X: a.X + b.X,
	}
}

var (
	north = Coord{1, 0}
	south = Coord{-1, 0}
	east  = Coord{0, 1}
	west  = Coord{0, -1}
)

type Graph struct {
	m map[Coord]byte
}

func (m Graph) Root() Coord {
	for key, value := range m.m {
		if value == 'S' {
			return key
		}
	}

	return Coord{-1, -1}
}

func (m Graph) Edges(c Coord) []Coord {
	connections := m.getConnections(c)

	edges := make([]Coord, 0, len(connections))

	for i := range connections {
		edge := add(c, connections[i])
		value := m.m[edge]

		switch connections[i] {
		case north:
			switch value {
			case '|', '7', 'F':
				edges = append(edges, edge)
			}

		case south:
			switch value {
			case '|', 'L', 'J':
				edges = append(edges, edge)
			}

		case east:
			switch value {
			case '-', '7', 'J':
				edges = append(edges, edge)
			}

		case west:
			switch value {
			case '-', 'F', 'L':
				edges = append(edges, edge)
			}
		}
	}

	return edges
}

func (m Graph) getConnections(c Coord) []Coord {
	switch m.m[c] {
	case '|':
		return []Coord{north, south}
	case '-':
		return []Coord{east, west}
	case 'L':
		return []Coord{north, east}
	case 'J':
		return []Coord{north, west}
	case '7':
		return []Coord{south, west}
	case 'F':
		return []Coord{south, east}
	case 'S':
		return []Coord{north, south, east, west}
	default:
		return []Coord{}
	}
}

func (m Graph) IsLast(c Coord) bool {
	return c == m.Root()
}

func Parse(input string) Graph {
	if input == "" {
		return Graph{}
	}

	var (
		scanner = bufio.NewScanner(strings.NewReader(input))
		y       = -1
		m       = Graph{
			m: make(map[Coord]byte),
		}
	)

	for scanner.Scan() {
		y++
		line := scanner.Text()

		for x := range line {
			m.m[Coord{Y: -y, X: x}] = line[x]
		}
	}

	return m
}

func HalfwayDistance(m Graph) int {
	dfs := graph.NewDFS[Coord]()

	dfs.Run(m)
	var n int

	for _, v := range dfs.Len {
		if int(v) > n {
			n = int(v)
		}
	}

	return (n + 1) / 2
}
