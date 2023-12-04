package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day04"
)

func TestNewSet(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants []Scratchcard
	}{
		{
			name: "OneLine",
			input: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
`,
			wants: []Scratchcard{{
				id: 1,
				winning: map[int]struct{}{
					41: {}, 48: {}, 83: {}, 86: {}, 17: {},
				},
				input: []int{
					83, 86, 6, 31, 17, 9, 48, 53,
				},
			}},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			cards := NewSet(testcase.input)

			isEqual(t, len(testcase.wants), len(cards))

			for i := range testcase.wants {
				isEqual(t, testcase.wants[i].id, cards[i].id)

				isEqual(t, len(testcase.wants[i].winning), len(cards[i].winning))

				values := make([]int, 0, len(testcase.wants[i].winning))
				for k := range testcase.wants[i].winning {
					values = append(values, k)
				}

				for idx := range values {
					_, ok := cards[i].winning[values[idx]]
					isEqual(t, true, ok)
				}

				isEqual(t, len(testcase.wants[i].input), len(cards[i].input))
				for idx := range testcase.wants[i].input {
					isEqual(t, testcase.wants[i].input[idx], cards[i].input[idx])
				}
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
			input: day04.Input,
			wants: 26218,
		},
		{
			name: "Example",
			input: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`,
			wants: 13,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			cards := NewSet(testcase.input)
			result := Sum(cards...)

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
