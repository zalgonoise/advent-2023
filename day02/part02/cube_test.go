package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day02"
)

func TestCubeConundrum(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day02.Input,
			wants: 70924,
		},
		{
			name: "Example",
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`,
			wants: 2286,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			sum, err := CubeConundrum(testcase.input)
			if err != nil {
				t.Error(err)

				return
			}

			isEqual(t, testcase.wants, sum)
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
