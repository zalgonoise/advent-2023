package grid

import (
	"fmt"
	"testing"
)

func TestRebuild(t *testing.T) {
	for _, quadrant := range []Quadrant{Q1, Q2, Q3, Q4} {
		t.Run(fmt.Sprintf("%d", quadrant+1), func(t *testing.T) {
			for _, testcase := range []struct {
				name  string
				input [][]int
			}{
				{
					name: "Simple",
					input: [][]int{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 11, 12},
						{13, 14, 15, 16},
					},
				},
			} {
				{
					t.Run(testcase.name, func(t *testing.T) {
						grid := New(testcase.input, WithQuadrant(quadrant))

						rebuilt := Rebuild(grid)

						t.Log(rebuilt)

						isEqual(t, len(testcase.input), len(rebuilt))
						for i := range testcase.input {
							isEqual(t, len(testcase.input[i]), len(rebuilt[i]))
							for idx := range testcase.input[i] {
								isEqual(t, testcase.input[i][idx], rebuilt[i][idx])
							}
						}
					})
				}
			}
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
