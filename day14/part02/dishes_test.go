package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day14"
)

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		iter  int
		wants int
	}{
		{
			name:  "Input",
			input: day14.Input,
			iter:  1000000000,
			wants: 83790,
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
			iter:  1000000000,
			wants: 64,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			g := Parse(testcase.input)
			result := g.Sum(testcase.iter)

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
