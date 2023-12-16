package part01

import (
	"strings"
)

var (
	north = Coord{1, 0}
	east  = Coord{0, 1}
	south = Coord{-1, 0}
	west  = Coord{0, -1}
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

func Rotate(c Coord, ccw bool) Coord {
	if ccw {
		switch c {
		case north:
			return west
		case west:
			return south
		case south:
			return east
		case east:
			return north
		}

		return Coord{}
	}

	switch c {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	}

	return Coord{}
}

type Grid struct {
	maxY int
	maxX int
	grid map[Coord]byte
}

type Point struct {
	cur Coord
	dir Coord
}

func Parse(input string) *Grid {
	if input == "" {
		return nil
	}

	raw := strings.Split(input, "\n")
	lines := make([]string, 0, len(raw))
	m := make(map[Coord]byte)

	for y := range raw {
		if raw[y] == "" {
			continue
		}

		lines = append(lines, raw[y])
	}

	maxY := len(lines) - 1
	maxX := len(lines[0]) - 1

	for y := range lines {
		for x := range lines[y] {
			switch lines[y][x] {
			case '.', '-', '|', '/', '\\':
				m[Coord{maxY - y, x}] = lines[y][x]
			default:
			}
		}
	}

	return &Grid{
		maxY: maxY,
		maxX: maxX,
		grid: m,
	}
}

func Sum(g *Grid) int {
	cache := make(map[Coord]map[Coord]struct{})
	scan(g, cache, Point{
		cur: Coord{g.maxY, -1},
		dir: east,
	})

	return len(cache)
}

func scan(g *Grid, cache map[Coord]map[Coord]struct{}, point Point) {
	nextCoord := Add(point.cur, point.dir)
	next, ok := g.grid[nextCoord]
	if !ok {
		return
	}

	directions, ok := cache[nextCoord]
	switch ok {
	case true:
		if _, ok = directions[point.dir]; ok {
			return
		}
	default:
		cache[nextCoord] = make(map[Coord]struct{})
	}

	cache[nextCoord][point.dir] = struct{}{}

	switch next {
	case '.':
		scan(g, cache, Point{nextCoord, point.dir})

	case '-':
		switch point.dir {
		case north, south:
			scan(g, cache, Point{nextCoord, east})
			scan(g, cache, Point{nextCoord, west})
		default:
			scan(g, cache, Point{nextCoord, point.dir})
		}

	case '|':
		switch point.dir {
		case east, west:
			scan(g, cache, Point{nextCoord, north})
			scan(g, cache, Point{nextCoord, south})
		default:
			scan(g, cache, Point{nextCoord, point.dir})
		}

	case '/':
		switch point.dir {
		case east, west:
			dir := Rotate(point.dir, true)
			scan(g, cache, Point{nextCoord, dir})
		default:
			dir := Rotate(point.dir, false)
			scan(g, cache, Point{nextCoord, dir})
		}

	case '\\':
		switch point.dir {
		case north, south:
			dir := Rotate(point.dir, true)
			scan(g, cache, Point{nextCoord, dir})
		default:
			dir := Rotate(point.dir, false)
			scan(g, cache, Point{nextCoord, dir})
		}
	}
}
