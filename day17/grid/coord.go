package grid

var (
	North = Coord{1, 0}
	South = Coord{-1, 0}
	East  = Coord{0, 1}
	West  = Coord{0, -1}
)

func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

type Coord struct {
	Y int
	X int
}

func Add(a, b Coord) Coord {
	return Coord{
		Y: a.Y + b.Y,
		X: a.X + b.X,
	}
}

func Sub(a, b Coord) Coord {
	return Coord{
		Y: a.Y - b.Y,
		X: a.X - b.X,
	}
}

func Inverse(c Coord) Coord {
	return Coord{
		Y: -c.Y,
		X: -c.X,
	}
}

func Manhattan(a, b Coord) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}
