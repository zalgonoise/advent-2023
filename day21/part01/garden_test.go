package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day21"
)

func TestSolve(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		iter  int
		wants int
	}{
		{
			name:  "Input/Part1",
			input: day21.Input,
			iter:  64,
			wants: 3724,
		},
		{
			name:  "Input/Part2",
			input: day21.Input,
			iter:  26501365,
			wants: 7602,
		},
		{
			name:  "Example",
			input: day21.TestInput,
			iter:  6,
			wants: 16,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			g := Parse(testcase.input, testcase.iter)
			result := Count(g)

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
