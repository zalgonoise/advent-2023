package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day06"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants Race
	}{
		{
			name: "Example",
			input: `Time:      7  15   30
Distance:  9  40  200
`,
			wants: Race{
				71530, 940200,
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			race, err := Parse(testcase.input)

			isEqual(t, nil, err)
			isEqual(t, testcase.wants.total, race.total)
			isEqual(t, testcase.wants.record, race.record)
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
			wants: 23501589,
		},
		{
			name: "Example",
			input: `Time:      7  15   30
Distance:  9  40  200
`,
			wants: 71503,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			race, err := Parse(testcase.input)

			isEqual(t, nil, err)

			result := Plan(race)

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
