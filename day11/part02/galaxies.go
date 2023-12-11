package part02

import "strings"

type Coord struct {
	Y int
	X int
}

type Grid struct {
	items    []Coord
	itemRows map[int]struct{}
	itemCols map[int]struct{}
}

func Parse(input string) *Grid {
	if input == "" {
		return nil
	}

	lines := strings.Split(input, "\n")
	grid := &Grid{
		items:    make([]Coord, 0, len(lines)*2),
		itemRows: make(map[int]struct{}, len(lines)),
		itemCols: make(map[int]struct{}, len(lines[0])),
	}

	for y := range lines {
		if lines[y] == "" {
			continue
		}

		for x := range lines[y] {
			if lines[y][x] == '#' {
				grid.items = append(grid.items, Coord{y, x})
				grid.itemRows[y] = struct{}{}
				grid.itemCols[x] = struct{}{}
			}
		}
	}

	return grid
}

func Sum(g *Grid, factor int) int {
	pairs := pair(g.items)

	var n int
	for i := range pairs {
		var distance int

		for idx := min(pairs[i][0].X, pairs[i][1].X); idx < max(pairs[i][0].X, pairs[i][1].X); idx++ {
			_, ok := g.itemCols[idx]
			switch ok {
			case true:
				distance++
			default:
				distance += factor
			}
		}

		for idx := min(pairs[i][0].Y, pairs[i][1].Y); idx < max(pairs[i][0].Y, pairs[i][1].Y); idx++ {
			_, ok := g.itemRows[idx]
			switch ok {
			case true:
				distance++
			default:
				distance += factor
			}
		}

		n += distance
	}

	return n
}

func pair[T any, S ~[]T](items S) [][2]T {
	if len(items) < 2 {
		return nil
	}

	pairs := make([][2]T, len(items)*2)

	for i := range items {
		for idx := i + 1; idx < len(items); idx++ {
			pairs = append(pairs, [2]T{items[i], items[idx]})
		}
	}

	return pairs
}
