package part01

import (
	"github.com/zalgonoise/advent-2023/day02/parser"
)

func CubeConundrum(input string, red, green, blue int) (int, error) {
	games, err := parser.Parse(input)
	if err != nil {
		return -1, err
	}

	var n int

	for i := range games {
		if games[i].LessThan(red, green, blue) {
			n += games[i].ID
		}
	}

	return n, nil
}
