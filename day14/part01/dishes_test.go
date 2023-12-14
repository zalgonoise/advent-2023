package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day14"
)

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day14.Input,
			wants: 103614,
		},
		{
			name: "Example",
			input: `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`,
			wants: 136,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			g := Parse(testcase.input)
			g.Tilt()

			result := g.Sum()

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
