package part02

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/zalgonoise/advent-2023/day12"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name   string
		input  string
		wants  []Set
		factor int
	}{
		{
			name: "Example/2",
			input: `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`,
			factor: 5,
			wants: []Set{
				{items: "???.###????.###????.###????.###????.###", combos: []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}},
				{items: ".??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.", combos: []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}},
				{items: "?#?#?#?#?#?#?#???#?#?#?#?#?#?#???#?#?#?#?#?#?#???#?#?#?#?#?#?#???#?#?#?#?#?#?#?", combos: []int{1, 3, 1, 6, 1, 3, 1, 6, 1, 3, 1, 6, 1, 3, 1, 6, 1, 3, 1, 6}},
				{items: "????.#...#...?????.#...#...?????.#...#...?????.#...#...?????.#...#...", combos: []int{4, 1, 1, 4, 1, 1, 4, 1, 1, 4, 1, 1, 4, 1, 1}},
				{items: "????.######..#####.?????.######..#####.?????.######..#####.?????.######..#####.?????.######..#####.", combos: []int{1, 6, 5, 1, 6, 5, 1, 6, 5, 1, 6, 5, 1, 6, 5}},
				{items: "?###??????????###??????????###??????????###??????????###????????", combos: []int{3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1}},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			sets, err := Parse(testcase.input, testcase.factor)

			isEqual(t, nil, err)

			isEqual(t, len(testcase.wants), len(sets))
			for i := range testcase.wants {
				isEqual(t, testcase.wants[i].items, sets[i].items, i, "items")
				isEqual(t, true, slices.Equal(testcase.wants[i].combos, sets[i].combos), i, "combos")
			}
		})
	}
}

func TestSum(t *testing.T) {
	for _, testcase := range []struct {
		name   string
		input  string
		factor int
		wants  int
	}{
		{
			name:   "Input",
			input:  day12.Input,
			factor: 5,
			wants:  37366887898686,
		},
		{
			name:   "Input/Part01",
			input:  day12.Input,
			factor: 1,
			wants:  7916,
		},
		{
			name: "Example",
			input: `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`,
			factor: 5,
			wants:  525152,
		},
		{
			name:   "Example/1",
			input:  `???.### 1,1,3`,
			factor: 5,
			wants:  1,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			sets, err := Parse(testcase.input, testcase.factor)

			isEqual(t, nil, err)

			result := Sum(sets)

			isEqual(t, testcase.wants, result)
		})
	}
}

func isEqual[T comparable](t *testing.T, wants, got T, args ...any) {
	if got != wants {
		str := fmt.Sprintf("output mismatch error: wanted %v ; got %v", wants, got)

		if len(args) > 0 {
			sb := &strings.Builder{}

			for i := range args {
				sb.WriteString(fmt.Sprintf(" -- %v", args[i]))
			}

			str = fmt.Sprint(str, sb.String())
		}

		t.Error(str)

		t.Fail()

		return
	}

	str := fmt.Sprintf("output matched expected value: %v", wants)

	if len(args) > 0 {
		sb := &strings.Builder{}

		for i := range args {
			sb.WriteString(fmt.Sprintf(" -- %v", args[i]))
		}

		str = fmt.Sprint(str, sb.String())
	}

	t.Logf(str)
}
