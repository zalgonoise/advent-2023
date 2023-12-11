package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day11"
)

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day11.Input,
			wants: 9521776,
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
			wants: 374,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			grid := Parse(testcase.input)

			result := Sum(grid, 2)

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
