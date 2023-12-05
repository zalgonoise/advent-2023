package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day05"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name       string
		input      string
		wantsSeeds []int
		wantsPlots []Plot
	}{
		{
			name: "Simple",
			input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48`,
			wantsSeeds: []int{79, 14, 55, 13},
			wantsPlots: []Plot{{
				{50, 98, 2}, {52, 50, 48},
			}},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			seeds, plots, err := Parse(testcase.input)

			isEqual(t, nil, err)
			isEqual(t, len(testcase.wantsSeeds), len(seeds))
			for i := range testcase.wantsSeeds {
				isEqual(t, testcase.wantsSeeds[i], seeds[i])
			}

			isEqual(t, len(testcase.wantsPlots), len(plots))
			for i := range testcase.wantsPlots {
				isEqual(t, len(testcase.wantsPlots[i]), len(plots[i]))

				for idx := range testcase.wantsPlots[i] {
					isEqual(t, testcase.wantsPlots[i][idx].from, plots[i][idx].from)
					isEqual(t, testcase.wantsPlots[i][idx].to, plots[i][idx].to)
					isEqual(t, testcase.wantsPlots[i][idx].span, plots[i][idx].span)
				}
			}
		})
	}
}

func TestLowest(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day05.Input,
			wants: 322500873,
		},
		{
			name: "Example",
			input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			wants: 35,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			seeds, plots, err := Parse(testcase.input)

			isEqual(t, nil, err)

			result := Lowest(seeds, plots...)

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
