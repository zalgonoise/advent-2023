package part02

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errMalformedInput = errors.New("malformed input; needs to contain two separate sections for moves and parts")
	errZeroIdxKey     = errors.New("key part of the move is empty")
)

type Operation int

const (
	noOp Operation = iota
	greaterThan
	lessThan
)

type Part struct {
	X, M, A, S int
}

type Condition struct {
	Op     Operation
	Target string
	Value  int
	Result Outcome[string]
}

type Outcome[T any] struct {
	Done bool
	Drop bool
	Key  T
}

type Pipeline struct {
	Moves map[string][]Condition
}

type Range[T any] struct {
	Min T
	Max T
}

func Parse(input string) (Pipeline, error) {
	if input == "" {
		return Pipeline{}, nil
	}

	movesAndParts := strings.Split(input, "\n\n")
	if len(movesAndParts) != 2 {
		return Pipeline{}, errMalformedInput
	}

	return extractMoves(movesAndParts[0])
}

func (p Pipeline) Travel(r Range[Part]) int {
	return p.travel("in", r)
}

func (p Pipeline) travel(start string, r Range[Part]) int {
	var n int
	conditions := p.Moves[start]

	for _, i := range conditions {
		bounds := i.split(r)
		switch {
		case i.Result.Done:
			n += sum(bounds[0])
		case i.Result.Drop, i.Result.Key == "":
		default:
			n += p.travel(i.Result.Key, bounds[0])
		}

		r = bounds[1]
	}

	return n
}

func (c Condition) split(r Range[Part]) [2]Range[Part] {
	inner := r
	outer := r

	switch c.Target {
	case "x":
		minVal, maxVal := r.Min.X, r.Max.X
		inner.Min.X, inner.Max.X = reduce(c.Op, c.Value, minVal, maxVal)
		outer.Min.X, outer.Max.X = invReduce(c.Op, c.Value, minVal, maxVal)

	case "m":
		minVal, maxVal := r.Min.M, r.Max.M
		inner.Min.M, inner.Max.M = reduce(c.Op, c.Value, minVal, maxVal)
		outer.Min.M, outer.Max.M = invReduce(c.Op, c.Value, minVal, maxVal)

	case "a":
		minVal, maxVal := r.Min.A, r.Max.A
		inner.Min.A, inner.Max.A = reduce(c.Op, c.Value, minVal, maxVal)
		outer.Min.A, outer.Max.A = invReduce(c.Op, c.Value, minVal, maxVal)

	case "s":
		minVal, maxVal := r.Min.S, r.Max.S
		inner.Min.S, inner.Max.S = reduce(c.Op, c.Value, minVal, maxVal)
		outer.Min.S, outer.Max.S = invReduce(c.Op, c.Value, minVal, maxVal)
	}

	return [2]Range[Part]{inner, outer}
}

func (p Part) valueFor(a string) (int, bool) {
	switch a {
	case "x":
		return p.X, true
	case "m":
		return p.M, true
	case "a":
		return p.A, true
	case "s":
		return p.S, true
	default:
		return 0, false
	}
}

func sum(r Range[Part]) int {
	switch {
	case r.Max.X <= r.Min.X,
		r.Max.M <= r.Min.M,
		r.Max.A <= r.Min.A,
		r.Max.S <= r.Min.S:
		return 0
	default:
		n := r.Max.X - r.Min.X + 1
		n *= r.Max.M - r.Min.M + 1
		n *= r.Max.A - r.Min.A + 1
		n *= r.Max.S - r.Min.S + 1

		return n
	}
}

func reduce(op Operation, value, minimum, maximum int) (int, int) {
	switch op {
	case noOp:
		return minimum, maximum
	case lessThan:
		if maximum >= value {
			maximum = value - 1
		}
	case greaterThan:
		if minimum <= value {
			minimum = value + 1
		}
	}

	return minimum, maximum
}

func invReduce(op Operation, value, minimum, maximum int) (int, int) {
	switch op {
	case noOp:
		return 0, 0
	case lessThan:
		if minimum < value {
			minimum = value
		}
	case greaterThan:
		if maximum > value {
			maximum = value
		}
	}

	return minimum, maximum
}

func extractMoves(input string) (Pipeline, error) {
	raw := strings.Split(input, "\n")
	lines := make([]string, 0, len(raw))

	for i := range raw {
		if raw[i] == "" {
			continue
		}

		lines = append(lines, raw[i])
	}

	p := Pipeline{
		Moves: make(map[string][]Condition, len(lines)),
	}

	for i := range lines {
		key, states, err := extractMove(lines[i])
		if err != nil {
			return Pipeline{}, err
		}

		p.Moves[key] = states
	}

	return p, nil
}

func extractMove(input string) (string, []Condition, error) {
	start := strings.Index(input, "{")
	end := strings.Index(input, "}")
	key := input[:start]
	raw := input[start+1 : end]

	split := strings.Split(raw, ",")
	states := make([]Condition, 0, len(split))

	for i := range split {
		state, err := extractState(split[i])
		if err != nil {
			return "", nil, err
		}

		states = append(states, state)
	}

	return key, states, nil
}

func extractState(input string) (Condition, error) {
	var (
		idx    int
		isDone bool
		toDrop bool
	)

	// first pass for the key (or outcome)
	for ; idx < len(input); idx++ {
		if input[idx] >= 'a' && input[idx] <= 'z' {
			continue
		}

		if input[idx] == 'A' {
			isDone = true
		}

		if input[idx] == 'R' {
			toDrop = true
		}

		break
	}

	// handle 'A'-only moves
	if isDone {
		return Condition{
			Target: "",
			Op:     noOp,
			Value:  0,
			Result: Outcome[string]{Done: true},
		}, nil
	}

	// handle 'R'-only moves
	if toDrop {
		return Condition{
			Target: "",
			Op:     noOp,
			Value:  0,
			Result: Outcome[string]{Drop: true},
		}, nil
	}

	if idx <= 0 {
		return Condition{}, errZeroIdxKey
	}

	var gt bool
	k := input[:idx]

	if idx == len(input) {
		return Condition{
			Target: "",
			Op:     noOp,
			Value:  0,
			Result: Outcome[string]{Key: k},
		}, nil
	}

	// grab the comparison type
	switch input[idx] {
	case '>':
		gt = true
	case '<':
		gt = false
	}

	idx++
	from := idx

	// scan the comparison value
	for ; idx < len(input); idx++ {
		if input[idx] >= '0' && input[idx] <= '9' {
			continue
		}

		break
	}

	num, err := strconv.Atoi(input[from:idx])
	if err != nil {
		return Condition{}, err
	}

	// handle travel / reject returns
	idx++
	switch input[idx] {
	case 'A':
		isDone = true
	case 'R':
		toDrop = true
	}

	// if key matches comparison >> accepted
	if isDone {
		switch gt {
		case true:
			return Condition{
				Target: k,
				Op:     greaterThan,
				Value:  num,
				Result: Outcome[string]{Done: true},
			}, nil
		default:
			return Condition{
				Target: k,
				Op:     lessThan,
				Value:  num,
				Result: Outcome[string]{Done: true},
			}, nil
		}
	}

	// if key matches comparison >> rejected
	if toDrop {
		switch gt {
		case true:
			return Condition{
				Target: k,
				Op:     greaterThan,
				Value:  num,
				Result: Outcome[string]{Drop: true},
			}, nil
		default:
			return Condition{
				Target: k,
				Op:     lessThan,
				Value:  num,
				Result: Outcome[string]{Drop: true},
			}, nil
		}
	}

	// if key matches comparison >> sent to next mapping reference
	to := input[idx:]

	switch gt {
	case true:
		return Condition{
			Target: k,
			Op:     greaterThan,
			Value:  num,
			Result: Outcome[string]{Key: to},
		}, nil
	default:
		return Condition{
			Target: k,
			Op:     lessThan,
			Value:  num,
			Result: Outcome[string]{Key: to},
		}, nil
	}
}
