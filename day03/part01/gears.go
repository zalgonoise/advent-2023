package part01

import (
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	minAlloc  = 16
	numberSet = "0123456789"
)

var (
	errYAxisOverflow = errors.New("overflow in the Y-axis")
	errXAxisOverflow = errors.New("overflow in the X-axis")
)

type Matrix struct {
	values  [][]byte
	symbols []Coordinate
	numbers []CoordRange
}

type Coordinate struct {
	X int
	Y int
}

type CoordRange struct {
	XStart int
	XEnd   int
	Y      int
}

func (m *Matrix) NearSymbol(distance int) []CoordRange {
	nearby := make([]CoordRange, 0, minAlloc)

	for idx := range m.symbols {
		for i := range m.numbers {
			valid := m.numbers[i]
			// check Y axis first
			if valid.Y-distance > m.symbols[idx].Y || valid.Y+distance < m.symbols[idx].Y {
				continue
			}

			if valid.XStart-distance <= m.symbols[idx].X && valid.XEnd-1+distance >= m.symbols[idx].X {
				nearby = append(nearby, m.numbers[i])
			}
		}
	}

	return nearby
}

func compare(a, b CoordRange) int {
	c := cmp.Compare(a.Y, b.Y)

	if c != 0 {
		return c
	}

	return cmp.Compare(a.XStart, b.XStart)
}

func (m *Matrix) Sum(ranges []CoordRange) (int, error) {
	errs := make([]error, 0, len(ranges))
	var n int

	for i := range ranges {
		value, err := m.rangeAsInt(ranges[i])
		if err != nil {
			errs = append(errs, err)

			continue
		}

		n += value
	}

	if len(errs) > 0 {
		return -1, errors.Join(errs...)
	}

	return n, nil
}

func (m *Matrix) rangeAsInt(r CoordRange) (int, error) {
	if r.Y >= len(m.values) {
		return -1, fmt.Errorf("%w: %d:%d", errYAxisOverflow, r.XStart, r.Y)
	}

	if r.XStart >= len(m.values[r.Y]) {
		return -1, fmt.Errorf("%w: %d:%d", errXAxisOverflow, r.XStart, r.Y)
	}

	if r.XEnd > len(m.values[r.Y]) {
		return -1, fmt.Errorf("%w: %d:%d", errXAxisOverflow, r.XEnd, r.Y)
	}

	n, err := strconv.Atoi(string(m.values[r.Y][r.XStart:r.XEnd]))
	if err != nil {
		return -1, err
	}

	return n, nil
}

func symbols(values [][]byte) []Coordinate {
	coords := make([]Coordinate, 0, minAlloc)

	for row := range values {
		for col, value := range values[row] {
			if (value >= '!' && value <= '-') ||
				value == '/' ||
				(value >= ':' && value <= '@') ||
				(value >= '[' && value <= '`') ||
				(value >= '{' && value <= '~') {
				coords = append(coords, Coordinate{col, row})
			}
		}
	}

	return coords
}

func numbers(values [][]byte) []CoordRange {
	coords := make([]CoordRange, 0, minAlloc)

	for row, value := range values {
		var (
			n     int
			start int
			end   int
		)

		if len(value) == 0 {
			continue
		}

		for len(value) > 0 {
			start = bytes.IndexAny(value, numberSet)

			if start == -1 {
				break
			}

			end = bytes.IndexFunc(value[start:], func(r rune) bool {
				return !unicode.IsNumber(r)
			})

			if end == -1 {
				end = len(value[start:])
			}

			value = value[start+end:]

			coords = append(coords, CoordRange{n + start, n + start + end, row})
			n += start + end
		}
	}

	return coords
}

func NewMatrix(input string) *Matrix {
	if input == "" {
		return nil
	}

	split := strings.Split(input, "\n")

	values := make([][]byte, 0, len(split))
	for i := range split {
		values = append(values, []byte(split[i]))
	}

	return &Matrix{
		values:  values,
		symbols: symbols(values),
		numbers: numbers(values),
	}
}
