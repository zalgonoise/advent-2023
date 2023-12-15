package part02

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day15"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants []Op
	}{
		{
			name:  "Example",
			input: "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			wants: []Op{
				{key: "rn", hash: Hash("rn"), op: int('='), value: 1},
				{key: "cm", hash: Hash("cm"), op: int('-')},
				{key: "qp", hash: Hash("qp"), op: int('='), value: 3},
				{key: "cm", hash: Hash("cm"), op: int('='), value: 2},
				{key: "qp", hash: Hash("qp"), op: int('-')},
				{key: "pc", hash: Hash("pc"), op: int('='), value: 4},
				{key: "ot", hash: Hash("ot"), op: int('='), value: 9},
				{key: "ab", hash: Hash("ab"), op: int('='), value: 5},
				{key: "pc", hash: Hash("pc"), op: int('-')},
				{key: "pc", hash: Hash("pc"), op: int('='), value: 6},
				{key: "ot", hash: Hash("ot"), op: int('='), value: 7},
			},
		},
	} {
		result := Parse(testcase.input)

		isEqual(t, len(testcase.wants), len(result))
		for i := range testcase.wants {
			isEqual(t, testcase.wants[i].key, result[i].key)
			isEqual(t, testcase.wants[i].hash, result[i].hash)
			isEqual(t, testcase.wants[i].op, result[i].op)
			isEqual(t, testcase.wants[i].value, result[i].value)
		}
	}
}

func TestMap(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants map[int][]Op
	}{
		{
			name:  "Example",
			input: "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			wants: map[int][]Op{
				0: {
					{key: "rn", hash: Hash("rn"), op: int('='), value: 1},
					{key: "cm", hash: Hash("cm"), op: int('='), value: 2},
				},
				1: {},
				3: {
					{key: "ot", hash: Hash("ot"), op: int('='), value: 7},
					{key: "ab", hash: Hash("ab"), op: int('='), value: 5},
					{key: "pc", hash: Hash("pc"), op: int('='), value: 6},
				},
			},
		},
	} {
		ops := Parse(testcase.input)

		result := Map(ops...)

		isEqual(t, len(testcase.wants), len(result))
		for k, v := range testcase.wants {
			value, ok := result[k]

			isEqual(t, true, ok)
			isEqual(t, len(v), len(value))

			for i := range v {
				isEqual(t, v[i].key, value[i].key)
				isEqual(t, v[i].hash, value[i].hash)
				isEqual(t, v[i].op, value[i].op)
				isEqual(t, v[i].value, value[i].value)
			}
		}
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
			input: day15.Input,
			wants: 271384,
		},
		{
			name:  "Example",
			input: "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			wants: 145,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			steps := Parse(testcase.input)
			m := Map(steps...)
			result := Sum(m)

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
