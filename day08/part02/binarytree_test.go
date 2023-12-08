package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day08"
)

func TestFind(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day08.Input,
			wants: 15690466351717,
		},
		{
			name: "Example",
			input: `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`,
			wants: 6,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			moves, tree := Parse(testcase.input)
			result, err := Find("A", tree, moves)

			isEqual(t, nil, err)
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
