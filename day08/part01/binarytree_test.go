package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day08"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name       string
		input      string
		wantsMoves []int
		wantsMap   map[string][2]string
	}{
		{
			name: "Example",
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`,
			wantsMoves: []int{1, 0},
			wantsMap: map[string][2]string{
				"AAA": {"BBB", "CCC"},
				"BBB": {"DDD", "EEE"},
				"CCC": {"ZZZ", "GGG"},
				"DDD": {"DDD", "DDD"},
				"EEE": {"EEE", "EEE"},
				"GGG": {"GGG", "GGG"},
				"ZZZ": {"ZZZ", "ZZZ"},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			moves, tree := Parse(testcase.input)

			isEqual(t, len(testcase.wantsMoves), len(moves))
			for i := range testcase.wantsMoves {
				isEqual(t, testcase.wantsMoves[i], moves[i])
			}

			isEqual(t, len(testcase.wantsMap), len(tree))
			for k, v := range testcase.wantsMap {
				_, ok := tree[k]
				isEqual(t, true, ok)
				isEqual(t, v[0], tree[k][0])
				isEqual(t, v[1], tree[k][1])
			}
		})
	}
}

func TestFind(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day08.Input,
			wants: 20659,
		},
		{
			name: "Example/1",
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`,
			wants: 2,
		},
		{
			name: "Example/2",
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`,
			wants: 6,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			moves, tree := Parse(testcase.input)
			result, err := Find("AAA", tree, moves)

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
