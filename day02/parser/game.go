package parser

type Game struct {
	ID    int
	Cubes []Cubes
}

func (g Game) LessThan(red, green, blue int) bool {
	for i := range g.Cubes {
		if !g.Cubes[i].LessThan(red, green, blue) {
			return false
		}
	}

	return true
}

func (g Game) Min() (red, green, blue int) {
	for i := range g.Cubes {
		if g.Cubes[i].Red > red {
			red = g.Cubes[i].Red
		}

		if g.Cubes[i].Green > green {
			green = g.Cubes[i].Green
		}

		if g.Cubes[i].Blue > blue {
			blue = g.Cubes[i].Blue
		}
	}

	return red, green, blue
}

type Cubes struct {
	Red   int
	Green int
	Blue  int
}

func (c Cubes) LessThan(red, green, blue int) bool {
	if red < c.Red || green < c.Green || blue < c.Blue {
		return false
	}

	return true
}
