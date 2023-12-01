package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day01"
)

func TestTrebuchet(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input []string
		wants int
	}{
		{
			name:  "Challenge",
			input: []string{day01.Input},
			wants: 55488,
		},
		{
			name: "Example/AsString",
			input: []string{`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`},
			wants: 142,
		},
		{
			name:  "Example/AsStringSlice",
			input: []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
			wants: 142,
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
