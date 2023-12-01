package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day01"
)

func TestSingleTrebuchet(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "1",
			input: "two1nine",
			wants: 29,
		},

		{
			name:  "2",
			input: "eightwothree",
			wants: 83,
		},
		{
			name:  "3",
			input: "abcone2threexyz",
			wants: 13,
		},
		{
			name:  "4",
			input: "xtwone3four",
			wants: 24,
		},
		{
			name:  "5",
			input: "4nineeightseven2",
			wants: 42,
		},
		{
			name:  "6",
			input: "zoneight234",
			wants: 14,
		},
		{
			name:  "7",
			input: "7pqrstsixteen",
			wants: 76,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			n := trebuchet(testcase.input)

			isEqual(t, testcase.wants, n)
		})
	}
}

func TestTrebuchet(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input []string
		wants int
	}{
		{
			name:  "Challenge",
			input: []string{day01.Input},
			wants: 55614,
		},
		{
			name: "Example/AsString",
			input: []string{`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`},
			wants: 281,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			res := Trebuchet(testcase.input...)

			isEqual(t, testcase.wants, res)
		})
	}
}

func isEqual[T comparable](t *testing.T, wants, got T) {
	if got != wants {
		t.Errorf("output mismatch error: wanted %v ; got %v", wants, got)

		return
	}

	t.Logf("output matched expected value: %v", wants)
}
