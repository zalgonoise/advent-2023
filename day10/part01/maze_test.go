package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day10"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants Graph
	}{
		{
			name: "Example",
			input: `.....
.S-7.
.|.|.
.L-J.
.....
`,
			wants: Graph{
				m: map[Coord]byte{
					{0, 0}: '.', {0, 1}: '.', {0, 2}: '.', {0, 3}: '.', {0, 4}: '.',
					{-1, 0}: '.', {-1, 1}: 'S', {-1, 2}: '-', {-1, 3}: '7', {-1, 4}: '.',
					{-2, 0}: '.', {-2, 1}: '|', {-2, 2}: '.', {-2, 3}: '|', {-2, 4}: '.',
					{-3, 0}: '.', {-3, 1}: 'L', {-3, 2}: '-', {-3, 3}: 'J', {-3, 4}: '.',
					{-4, 0}: '.', {-4, 1}: '.', {-4, 2}: '.', {-4, 3}: '.', {-4, 4}: '.',
				},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			m := Parse(testcase.input)

			isEqual(t, len(testcase.wants.m), len(m.m))
			for k, v := range testcase.wants.m {
				mv, ok := m.m[k]

				isEqual(t, true, ok)
				isEqual(t, v, mv)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day10.Input,
			wants: 6649,
		},
		{
			name: "Example/1",
			input: `.....
.S-7.
.|.|.
.L-J.
.....
`,
			wants: 4,
		},
		{
			name: "Example/2",
			input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`,
			wants: 8,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			maze := Parse(testcase.input)

			result := HalfwayDistance(maze)

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
