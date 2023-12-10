package part02

import (
	"bufio"
	"strings"

	"github.com/zalgonoise/advent-2023/day10/graph"
)

type Coord struct {
	Y int
	X int
}

func Add(a, b Coord) Coord {
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
		edge := Add(c, connections[i])
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

func EnclosedTiles(m Graph) int {
	dfs := graph.NewDFS[Coord]()

	dfs.Run(m)

	for coord := range m.m {
		if _, ok := dfs.Cache[coord]; !ok {
			m.m[coord] = '.'
		}
	}

	var (
		root      = m.Root()
		edges     = m.Edges(root)
		rootNorth = Add(root, north)
		rootSouth = Add(root, south)
		rootEast  = Add(root, east)
		rootWest  = Add(root, west)
	)

	switch {
	case edges[0] == rootNorth && edges[1] == rootSouth:
		m.m[root] = '|'
	case edges[0] == rootNorth && edges[1] == rootEast:
		m.m[root] = 'L'
	case edges[0] == rootNorth && edges[1] == rootWest:
		m.m[root] = 'J'
	case edges[0] == rootSouth && edges[1] == rootEast:
		m.m[root] = 'F'
	case edges[0] == rootSouth && edges[1] == rootWest:
		m.m[root] = '7'
	case edges[0] == rootEast && edges[1] == rootWest:
		m.m[root] = '-'
	default:
		return -1
	}

	sum := 0

	for coord, value := range m.m {
		if value == '.' {
			count := 0
			for i := coord.Y + 1; i <= 0; i++ {
				nextValue := m.m[Coord{X: coord.X, Y: i}]

				switch nextValue {
				case '-', 'F', 'L':
					count++
				}
			}

			if count%2 != 0 {
				sum++
			}
		}
	}

	return sum
}
