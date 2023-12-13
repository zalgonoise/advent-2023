package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day13"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants []Field
	}{
		{
			name: "Example",
			input: `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`,
			wants: []Field{
				{rows: []string{
					"#.##..##.",
					"..#.##.#.",
					"##......#",
					"##......#",
					"..#.##.#.",
					"..##..##.",
					"#.#.##.#.",
				}},
				{
					rows: []string{
						"#...##..#",
						"#....#..#",
						"..##..###",
						"#####.##.",
						"#####.##.",
						"..##..###",
						"#....#..#",
					},
				},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			fields := Parse(testcase.input)

			isEqual(t, len(testcase.wants), len(fields))
			for i := range testcase.wants {
				isEqual(t, len(testcase.wants[i].rows), len(fields[i].rows))

				for idx := range testcase.wants[i].rows {
					isEqual(t, testcase.wants[i].rows[idx], fields[i].rows[idx])
				}
			}
		})
	}
}

func TestFindY(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name: "Example/1",
			input: `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.
`,
			wants: 3,
		},
		{
			name: "Example/2",
			input: `
#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`,
			wants: 1,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			fields := Parse(testcase.input)

			isEqual(t, 1, len(fields))

			result := find(fields[0].rows)

			isEqual(t, testcase.wants, result)
		})
	}
}

func TestRotate(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants []string
	}{
		{
			name: "Example",
			input: `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.
`,
			wants: []string{
				"#.##..#",
				"..##...",
				"##..###",
				"#....#.",
				".#..#.#",
				".#..#.#",
				"#....#.",
				"##..###",
				"..##...",
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			fields := Parse(testcase.input)

			isEqual(t, 1, len(fields))

			result := rotate(fields[0].rows)

			isEqual(t, len(testcase.wants), len(result))
			for i := range testcase.wants {
				isEqual(t, testcase.wants[i], result[i])
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
			input: day13.Input,
			wants: 28235,
		},
		{
			name: "Example",
			input: `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`,
			wants: 400,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			fields := Parse(testcase.input)

			result := Sum(fields)

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
