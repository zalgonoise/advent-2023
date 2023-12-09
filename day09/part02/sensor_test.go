package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day09"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants [][]int
	}{
		{
			name: "Example",
			input: `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`,
			wants: [][]int{
				{0, 3, 6, 9, 12, 15},
				{1, 3, 6, 10, 15, 21},
				{10, 13, 16, 21, 30, 45},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			sets, err := Parse(testcase.input)

			isEqual(t, nil, err)

			isEqual(t, len(testcase.wants), len(sets))
			for i := range testcase.wants {
				isEqual(t, len(testcase.wants[i]), len(sets[i]))
				for idx := range testcase.wants[i] {
					isEqual(t, testcase.wants[i][idx], sets[i][idx])
				}
			}
		})
	}
}

func TestNextValue(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input []int
		wants int
	}{
		{
			name:  "Example/1",
			input: []int{0, 3, 6, 9, 12, 15},
			wants: -3,
		},
		{
			name:  "Example/2",
			input: []int{1, 3, 6, 10, 15, 21},
			wants: 0,
		},
		{
			name:  "Example/3",
			input: []int{10, 13, 16, 21, 30, 45},
			wants: 5,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			next := nextValue(testcase.input)

			isEqual(t, testcase.wants, next)
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
			input: day09.Input,
			wants: 884,
		},
		{
			name: "Example",
			input: `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`,
			wants: 2,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			sets, err := Parse(testcase.input)

			isEqual(t, nil, err)

			result := Sum(sets)
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
