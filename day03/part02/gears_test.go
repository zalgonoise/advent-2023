package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day03"
)

func TestNewMatrix(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants *Matrix
	}{
		{
			name:  "OneLine",
			input: "467..114..",
			wants: &Matrix{
				values:  [][]byte{[]byte("467..114..")},
				symbols: make([]Coordinate, 0, minAlloc),
				numbers: []CoordRange{
					{XStart: 0, XEnd: 3, Y: 0},
					{XStart: 5, XEnd: 8, Y: 0},
				},
			},
		},
		{
			name:  "EndingInNumber",
			input: "....926",
			wants: &Matrix{
				values:  [][]byte{[]byte("....926")},
				symbols: make([]Coordinate, 0, minAlloc),
				numbers: []CoordRange{
					{XStart: 4, XEnd: 7, Y: 0},
				},
			},
		},
		{
			name: "AllExampleInput",
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`,
			wants: &Matrix{
				values: [][]byte{
					[]byte("467..114.."),
					[]byte("...*......"),
					[]byte("..35..633."),
					[]byte("......#..."),
					[]byte("617*......"),
					[]byte(".....+.58."),
					[]byte("..592....."),
					[]byte("......755."),
					[]byte("...$.*...."),
					[]byte(".664.598.."),
				},
				symbols: []Coordinate{
					{X: 3, Y: 1},
					{X: 6, Y: 3},
					{X: 3, Y: 4},
					{X: 5, Y: 5},
					{X: 3, Y: 8},
					{X: 5, Y: 8},
				},
				numbers: []CoordRange{
					{XStart: 0, XEnd: 3, Y: 0},
					{XStart: 5, XEnd: 8, Y: 0},
					{XStart: 2, XEnd: 4, Y: 2},
					{XStart: 6, XEnd: 9, Y: 2},
					{XStart: 0, XEnd: 3, Y: 4},
					{XStart: 7, XEnd: 9, Y: 5},
					{XStart: 2, XEnd: 5, Y: 6},
					{XStart: 6, XEnd: 9, Y: 7},
					{XStart: 1, XEnd: 4, Y: 9},
					{XStart: 5, XEnd: 8, Y: 9},
				},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			m := NewMatrix(testcase.input)

			isEqual(t, len(testcase.wants.symbols), len(m.symbols))
			isEqual(t, len(testcase.wants.numbers), len(m.numbers))

			for i := range testcase.wants.symbols {
				isEqual(t, testcase.wants.symbols[i].X, m.symbols[i].X)
				isEqual(t, testcase.wants.symbols[i].Y, m.symbols[i].Y)
			}

			for i := range testcase.wants.numbers {
				isEqual(t, testcase.wants.numbers[i].XStart, m.numbers[i].XStart)
				isEqual(t, testcase.wants.numbers[i].XEnd, m.numbers[i].XEnd)
				isEqual(t, testcase.wants.numbers[i].Y, m.numbers[i].Y)
			}
		})
	}
}

func TestNearSymbol(t *testing.T) {
	for _, testcase := range []struct {
		name     string
		input    string
		distance int
		wants    []CoordRange
	}{
		{
			name: "OneMatch",
			input: `467..114..
...*......`,
			distance: 1,
			wants: []CoordRange{
				{XStart: 0, XEnd: 3, Y: 0},
			},
		},
		{
			name: "OneMatchOnEdge",
			input: `..467..114
.........*`,
			distance: 1,
			wants: []CoordRange{
				{XStart: 7, XEnd: 10, Y: 0},
			},
		},
		{
			name: "OneMatchOnEdgeOffset",
			input: `..467..114
......*...`,
			distance: 1,
			wants: []CoordRange{
				{XStart: 7, XEnd: 10, Y: 0},
			},
		},
		{
			name: "OneMatchOffset2",
			input: `.467..114.
.........*`,
			distance: 1,
			wants: []CoordRange{
				{XStart: 6, XEnd: 9, Y: 0},
			},
		},
		{
			name: "FromInput",
			input: `............677..........................................................................227.....730..35.......318...........92...166.......
....%..863..#......................36.............956..337%......692..............*744....$..........*......../.....187..-..................
`,
			distance: 1,
			wants: []CoordRange{
				{XStart: 12, XEnd: 15, Y: 0},
				{XStart: 55, XEnd: 58, Y: 1},
				{XStart: 83, XEnd: 86, Y: 1},
				{XStart: 89, XEnd: 92, Y: 0},
				{XStart: 102, XEnd: 104, Y: 0},
				{XStart: 111, XEnd: 114, Y: 0},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			m := NewMatrix(testcase.input)
			ranges := m.NearSymbol(testcase.distance)

			isEqual(t, len(testcase.wants), len(ranges))

			for i := range testcase.wants {
				isEqual(t, testcase.wants[i].Y, ranges[i].Y)
				isEqual(t, testcase.wants[i].XStart, ranges[i].XStart)
				isEqual(t, testcase.wants[i].XEnd, ranges[i].XEnd)
			}
		})
	}
}

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name     string
		input    string
		distance int
		wants    int
	}{
		{
			name:     "Input",
			input:    day03.Input,
			distance: 1,
			wants:    72246648,
		},
		{
			name: "AllExampleInput",
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`,
			distance: 1,
			wants:    467835,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			m := NewMatrix(testcase.input)
			ranges := m.NearSymbol(testcase.distance)
			result, err := m.Sum(ranges)

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

func TestInputCharMap(t *testing.T) {
	m := make(map[string]int, 255)

	for i := range day03.Input {
		m[string(day03.Input[i])]++
	}

	t.Log(m)
}
