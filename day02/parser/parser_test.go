package parser

import "testing"

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants []Game
	}{
		{
			name: "OneEntry/1",
			input: `Game 1: 3 Blue, 4 Red; 1 Red, 2 Green, 6 Blue; 2 Green
`,
			wants: []Game{
				{ID: 1, Cubes: []Cubes{{Blue: 3, Red: 4}, {Red: 1, Green: 2, Blue: 6}, {Green: 2}}},
			},
		},
		{
			name: "OneEntry/2",
			input: `Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
`,
			wants: []Game{
				{ID: 2, Cubes: []Cubes{{Blue: 1, Green: 2}, {Green: 3, Blue: 4, Red: 1}, {Green: 1, Blue: 1}}},
			},
		},
		{
			name: "MultipleEntries",
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`,
			wants: []Game{
				{ID: 1, Cubes: []Cubes{{Blue: 3, Red: 4}, {Red: 1, Green: 2, Blue: 6}, {Green: 2}}},
				{ID: 2, Cubes: []Cubes{{Blue: 1, Green: 2}, {Green: 3, Blue: 4, Red: 1}, {Green: 1, Blue: 1}}},
				{ID: 3, Cubes: []Cubes{{Green: 8, Blue: 6, Red: 20}, {Blue: 5, Red: 4, Green: 13}, {Green: 5, Red: 1}}},
				{ID: 4, Cubes: []Cubes{{Green: 1, Red: 3, Blue: 6}, {Green: 3, Red: 6}, {Green: 3, Blue: 15, Red: 14}}},
				{ID: 5, Cubes: []Cubes{{Red: 6, Blue: 1, Green: 3}, {Blue: 2, Red: 1, Green: 2}}},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			games, err := Parse(testcase.input)

			if err != nil {
				t.Error(err)

				return
			}

			isEqual(t, len(testcase.wants), len(games))

			for i := range testcase.wants {
				isEqual(t, testcase.wants[i].ID, games[i].ID)

				isEqual(t, len(testcase.wants[i].Cubes), len(games[i].Cubes))

				for idx := range testcase.wants[i].Cubes {
					isEqual(t, testcase.wants[i].Cubes[idx].Red, games[i].Cubes[idx].Red)
					isEqual(t, testcase.wants[i].Cubes[idx].Green, games[i].Cubes[idx].Green)
					isEqual(t, testcase.wants[i].Cubes[idx].Blue, games[i].Cubes[idx].Blue)
				}
			}
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
