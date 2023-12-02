package part01

import (
	"github.com/zalgonoise/advent-2023/day02/parser"
)

func CubeConundrum(input string) (int, error) {
	games, err := parser.Parse(input)
	if err != nil {
		return -1, err
	}

	var n int

	for i := range games {
		r, g, b := games[i].Min()

		n += r * g * b
	}

	return n, nil
}
