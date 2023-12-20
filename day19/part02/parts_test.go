package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day19"
)

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		from  int
		to    int
		wants int
	}{
		{
			name:  "Input",
			input: day19.Input,
			from:  1,
			to:    4000,
			wants: 128163929109524,
		},
		{
			name:  "Example",
			input: day19.TestInput,
			from:  1,
			to:    4000,
			wants: 167409079868000,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			pipeline, err := Parse(testcase.input)

			isEqual(t, nil, err)

			r := Range[Part]{
				Min: Part{testcase.from, testcase.from, testcase.from, testcase.from},
				Max: Part{testcase.to, testcase.to, testcase.to, testcase.to},
			}

			result := pipeline.Travel(r)

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
