package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day19"
)

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day19.Input,
			wants: 0,
		},
		{
			name:  "Example",
			input: day19.TestInput,
			wants: 19114,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			pipeline, parts, err := Parse(testcase.input)

			isEqual(t, nil, err)

			outcome := pipeline.Travel(parts)

			result := Sum(outcome)

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
