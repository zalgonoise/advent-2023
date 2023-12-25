package part02

import (
	"strings"

	"github.com/zalgonoise/advent-2023/day21/graph"
)

func Parse(input string, steps int) *Graph {
	lines := readInput(input)

	grid := graph.NewGrid(lines, graph.WithQuadrant(graph.Q3))

	return newGraph(grid, steps)
}

func Count(g *Graph) int {
	g.bfs.Run(g)

	var n int
	for _, value := range g.bfs.Len {
		if value%2 == 0 {
			n++
		}
	}

	return n
}

func newGraph(m graph.Map[byte], steps int) *Graph {
	root := graph.Coord{-1, -1}
	for key, value := range m.Items {
		if value == 'S' {
			root = key

			break
		}
	}

	g := &Graph{
		steps:     steps,
		root:      root,
		positions: m.Items,
		bfs:       graph.NewBFS[graph.Coord](),
	}

	return g
}

func readInput(input string) [][]byte {
	raw := strings.Split(input, "\n")
	lines := make([][]byte, 0, len(raw))

	for i := range raw {
		if raw[i] == "" {
			continue
		}

		lines = append(lines, []byte(raw[i]))
	}

	return lines
}

type Graph struct {
	steps int
	root  graph.Coord

	positions map[graph.Coord]byte
	bfs       *graph.BFS[graph.Coord]
}

func (m Graph) Root() graph.Coord {
	return m.root
}

func (m Graph) Edges(pos graph.Coord) []graph.Coord {
	edges := make([]graph.Coord, 0, 4)

	if m.bfs.Len[pos] < uint64(m.steps) {
		for _, dir := range graph.Directions {
			next := graph.Add(pos, dir)

			if m.positions[next] == '.' {
				edges = append(edges, next)
			}
		}
	}

	return edges
}

func (m Graph) IsLast(pos graph.Coord) bool {
	return pos == m.root
}
