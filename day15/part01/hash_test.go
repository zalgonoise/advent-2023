package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day15"
)

func TestHash(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Example",
			input: "HASH",
			wants: 52,
		},
	} {
		result := Hash(testcase.input)

		isEqual(t, testcase.wants, result)
	}
}

func TestHashSum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day15.Input,
			wants: 506869,
		},
		{
			name:  "Example",
			input: "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			wants: 1320,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			steps := Parse(testcase.input)
			result := HashSum(steps...)

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
