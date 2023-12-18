package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day17"
)

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day17.Input,
			wants: 1023,
		},
		{
			name: "Example",
			input: `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
`,
			wants: 102,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			g := Parse(testcase.input)

			result := AStar(g)

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
