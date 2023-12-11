package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day11"
)

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name   string
		input  string
		factor int
		wants  int
	}{
		{
			name:   "Input",
			input:  day11.Input,
			factor: 1_000_000,
			wants:  553224415344,
		},
		{
			name: "Example",
			input: `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`,
			factor: 10,
			wants:  1030,
		},
		{
			name: "Example",
			input: `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`,
			factor: 100,
			wants:  8410,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			grid := Parse(testcase.input)

			result := Sum(grid, testcase.factor)

			isEqual(t, testcase.wants, result)
		})
	}
}

func isEqual[T comparable](t *testing.T, wants, got T) {
	if got != wants {
		t.Errorf("output mismatch error: wanted %v ; got %v", wants, got)
		t.Fail()

		return
	}

	t.Logf("output matched expected value: %v", wants)
}
