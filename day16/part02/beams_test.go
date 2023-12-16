package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day16"
)

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day16.Input,
			wants: 7716,
		},
		{
			name: "Example",
			input: `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
`,
			wants: 51,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			grid := Parse(testcase.input)

			result := Sum(grid)

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
