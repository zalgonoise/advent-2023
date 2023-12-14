package part01

import (
	"strings"
)

type Coord struct {
	Y int
	X int
}

type Grid struct {
	maxX  int
	maxY  int
	items map[Coord]byte
}

func Parse(input string) *Grid {
	lines := strings.Split(input, "\n")

	maxY := len(lines)
	maxX := len(lines[0])

	coords := make(map[Coord]byte)

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] != '.' {
				coords[Coord{X: x, Y: maxY - y - 1}] = lines[y][x]
			}
		}
	}

	return &Grid{maxX, maxY, coords}
}

func (g *Grid) Sum() int {
	var n int

	for coord, item := range g.items {
		if item == 'O' {
			n += coord.Y
		}
	}

	return n
}

func (g *Grid) Tilt() {
	for x := 0; x < g.maxX; x++ {
		stop := g.maxY

		for y := g.maxY - 1; y >= 0; y-- {
			curCoord := Coord{X: x, Y: y}
			cur, ok := g.items[curCoord]
			if !ok {
				continue
			}

			switch cur {
			case 'O':
				nextCoord := Coord{X: x, Y: stop - 1}

				g.items[curCoord] = 0
				g.items[nextCoord] = cur

				stop = stop - 1
			case '#':
				stop = y
			}
		}
	}
}
