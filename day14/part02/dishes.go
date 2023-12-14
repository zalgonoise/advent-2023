package part02

import (
	"fmt"
	"strings"
)

const minAlloc = 64

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
	raw := strings.Split(input, "\n")

	lines := make([]string, 0, len(raw))
	for i := range raw {
		if raw[i] != "" {
			lines = append(lines, raw[i])
		}
	}

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

func (g *Grid) Sum(iter int) int {
	cache := make(map[string]int)
	loop := true

	for i := 0; i < iter; i++ {
		g.Tilt()
		key := g.key()

		if idx, ok := cache[key]; ok && loop {
			interval := i - idx

			i = idx + ((iter-idx)/interval)*interval
			loop = false
		}

		cache[key] = i
	}

	var n int

	for coords, item := range g.items {
		if item == 'O' {
			n += coords.Y + 1
		}
	}
	return n
}

func (g *Grid) tiltNorth() {
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
				next := Coord{X: x, Y: stop - 1}

				delete(g.items, curCoord)
				g.items[next] = cur

				stop--
			case '#':
				stop = y
			}
		}
	}
}

func (g *Grid) tiltSouth() {
	for x := 0; x < g.maxX; x++ {
		stop := -1

		for y := 0; y < g.maxY; y++ {
			curCoord := Coord{X: x, Y: y}
			cur, ok := g.items[curCoord]
			if !ok {
				continue
			}

			switch cur {
			case 'O':
				next := Coord{X: x, Y: stop + 1}

				delete(g.items, curCoord)
				g.items[next] = cur

				stop++
			case '#':
				stop = y
			}
		}
	}
}

func (g *Grid) tiltEast() {
	for y := 0; y < g.maxY; y++ {
		stop := g.maxX

		for x := g.maxX - 1; x >= 0; x-- {
			curCoord := Coord{X: x, Y: y}
			cur, ok := g.items[curCoord]
			if !ok {
				continue
			}

			switch cur {
			case 'O':
				next := Coord{X: stop - 1, Y: y}
				delete(g.items, curCoord)
				g.items[next] = cur

				stop--
			case '#':
				stop = x
			}
		}
	}
}

func (g *Grid) tiltWest() {
	for y := 0; y < g.maxY; y++ {
		stop := -1

		for x := 0; x < g.maxX; x++ {
			curCoord := Coord{X: x, Y: y}
			cur, ok := g.items[curCoord]
			if !ok {
				continue
			}

			switch cur {
			case 'O':
				next := Coord{X: stop + 1, Y: y}

				delete(g.items, curCoord)
				g.items[next] = cur

				stop++
			case '#':
				stop = x
			}
		}
	}
}

func (g *Grid) Tilt() {
	g.tiltNorth()
	g.tiltWest()
	g.tiltSouth()
	g.tiltEast()
}

func (g *Grid) key() string {
	coords := make([]string, 0, minAlloc)

	for y := g.maxY - 1; y >= 0; y-- {
		for x := 0; x < g.maxX; x++ {
			coord := Coord{X: x, Y: y}
			cur, ok := g.items[coord]

			if ok && cur == 'O' {
				coords = append(coords, fmt.Sprintf("%d:%d", y, x))
			}
		}
	}

	return fmt.Sprintf("%d::%d::%s", g.maxY, g.maxX, strings.Join(coords, "::"))
}
