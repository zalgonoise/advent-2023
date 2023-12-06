package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day06"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants []Race
	}{
		{
			name: "Example",
			input: `Time:      7  15   30
Distance:  9  40  200
`,
			wants: []Race{
				{7, 9}, {15, 40}, {30, 200},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			races, err := Parse(testcase.input)

			isEqual(t, nil, err)
			isEqual(t, len(testcase.wants), len(races))

			for i := range testcase.wants {
				isEqual(t, testcase.wants[i].total, races[i].total)
				isEqual(t, testcase.wants[i].record, races[i].record)
			}
		})
	}
}

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day06.Input,
			wants: 1083852,
		},
		{
			name: "Example",
			input: `Time:      7  15   30
Distance:  9  40  200
`,
			wants: 288,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			races, err := Parse(testcase.input)

			isEqual(t, nil, err)

			result := Sum(races...)

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
