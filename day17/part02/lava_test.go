package part02

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
			wants: 1165,
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
