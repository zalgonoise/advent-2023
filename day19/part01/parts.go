package part01

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errMalformedInput     = errors.New("malformed input; needs to contain two separate sections for moves and parts")
	errZeroIdxKey         = errors.New("key part of the move is empty")
	errNotEnoughPartElems = errors.New("not enough part elements")
	errInvalidKeyValue    = errors.New("invalid key-value pair in part")
)

type Pipeline struct {
	Moves map[string][]State
}

func (p Pipeline) Travel(parts []Part) []Outcome[Part] {
	outcome := make([]Outcome[Part], 0, len(parts))

	for i := range parts {
		outcome = append(outcome, p.travel("in", parts[i]))
	}

	return outcome
}

func (p Pipeline) travel(key string, part Part) Outcome[Part] {
	states, ok := p.Moves[key]
	if !ok {
		return Outcome[Part]{
			Drop: true,
			Key:  part,
		}
	}

	for i := range states {
		if outcome, ok := states[i]("x", part.X); ok {
			return p.check(part, outcome)
		}

		if outcome, ok := states[i]("m", part.M); ok {
			return p.check(part, outcome)
		}

		if outcome, ok := states[i]("a", part.A); ok {
			return p.check(part, outcome)
		}

		if outcome, ok := states[i]("s", part.S); ok {
			return p.check(part, outcome)
		}
	}

	return Outcome[Part]{
		Drop: true,
		Key:  part,
	}
}

func (p Pipeline) check(part Part, outcome Outcome[string]) Outcome[Part] {
	switch {
	case outcome.Done:
		return Outcome[Part]{
			Done: true,
			Key:  part,
		}
	case outcome.Drop:
		return Outcome[Part]{
			Drop: true,
			Key:  part,
		}
	}

	return p.travel(outcome.Key, part)
}

func Sum(outcome []Outcome[Part]) int {
	var n int

	for i := range outcome {
		if outcome[i].Done {
			n += outcome[i].Key.X +
				outcome[i].Key.M +
				outcome[i].Key.A +
				outcome[i].Key.S
		}
	}

	return n
}

type State func(key string, value int) (Outcome[string], bool)

type Outcome[T any] struct {
	Done bool
	Drop bool
	Key  T
}

type Part struct {
	X, M, A, S int
}

func Parse(input string) (Pipeline, []Part, error) {
	if input == "" {
		return Pipeline{}, nil, nil
	}

	movesAndParts := strings.Split(input, "\n\n")
	if len(movesAndParts) != 2 {
		return Pipeline{}, nil, errMalformedInput
	}

	parts, err := extractParts(movesAndParts[1])
	if err != nil {
		return Pipeline{}, nil, err
	}

	pipeline, err := extractMoves(movesAndParts[0])
	if err != nil {
		return Pipeline{}, nil, err
	}

	return pipeline, parts, nil
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
		Moves: make(map[string][]State, len(lines)),
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

func extractMove(input string) (string, []State, error) {
	start := strings.Index(input, "{")
	end := strings.Index(input, "}")
	key := input[:start]
	raw := input[start+1 : end]

	split := strings.Split(raw, ",")
	states := make([]State, 0, len(split))

	for i := range split {
		state, err := extractState(split[i])
		if err != nil {
			return "", nil, err
		}

		states = append(states, state)
	}

	return key, states, nil
}

func extractState(input string) (State, error) {
	var (
		idx  int
		done bool
		drop bool
	)

	// first pass for the key (or outcome)
	for ; idx < len(input); idx++ {
		if input[idx] >= 'a' && input[idx] <= 'z' {
			continue
		}

		if input[idx] == 'A' {
			done = true
		}

		if input[idx] == 'R' {
			drop = true
		}

		break
	}

	// handle 'A'-only moves
	if done {
		return func(key string, value int) (Outcome[string], bool) {
			return Outcome[string]{Done: true}, true
		}, nil
	}

	// handle 'R'-only moves
	if drop {
		return func(key string, value int) (Outcome[string], bool) {
			return Outcome[string]{Drop: true}, true
		}, nil
	}

	if idx <= 0 {
		return nil, errZeroIdxKey
	}

	var greaterThan bool
	k := input[:idx]

	if idx == len(input) {
		return func(key string, value int) (Outcome[string], bool) {
			return Outcome[string]{Key: k}, true
		}, nil
	}

	// grab the comparison type
	switch input[idx] {
	case '>':
		greaterThan = true
	case '<':
		greaterThan = false
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
		return nil, err
	}

	// handle accept / reject returns
	idx++
	switch input[idx] {
	case 'A':
		done = true
	case 'R':
		drop = true
	}

	// if key matches comparison >> accepted
	if done {
		return func(key string, value int) (Outcome[string], bool) {
			if key != k {
				return Outcome[string]{}, false
			}

			switch greaterThan {
			case true:
				if value > num {
					return Outcome[string]{Done: true}, true
				}
			default:
				if value < num {
					return Outcome[string]{Done: true}, true
				}
			}

			return Outcome[string]{}, false
		}, nil
	}

	// if key matches comparison >> rejected
	if drop {
		return func(key string, value int) (Outcome[string], bool) {
			if key != k {
				return Outcome[string]{}, false
			}

			switch greaterThan {
			case true:
				if value > num {
					return Outcome[string]{Drop: true}, true
				}
			default:
				if value < num {
					return Outcome[string]{Drop: true}, true
				}
			}

			return Outcome[string]{}, false
		}, nil
	}

	// if key matches comparison >> sent to next mapping reference
	to := input[idx:]
	return func(key string, value int) (Outcome[string], bool) {
		if key != k {
			return Outcome[string]{}, false
		}

		switch greaterThan {
		case true:
			if value > num {
				return Outcome[string]{Key: to}, true
			}
		default:
			if value < num {
				return Outcome[string]{Key: to}, true
			}
		}

		return Outcome[string]{}, false
	}, nil
}

func extractParts(input string) ([]Part, error) {
	raw := strings.Split(input, "\n")
	lines := make([]string, 0, len(raw))

	for i := range raw {
		if raw[i] == "" {
			continue
		}

		lines = append(lines, raw[i])
	}

	parts := make([]Part, 0, len(lines))

	for i := range lines {
		part, err := extractPart(lines[i])
		if err != nil {
			return nil, err
		}

		parts = append(parts, part)
	}

	return parts, nil
}

func extractPart(input string) (Part, error) {
	part := Part{}
	start := strings.Index(input, "{")
	end := strings.Index(input, "}")
	input = input[start+1 : end]

	split := strings.Split(input, ",")
	if len(split) != 4 {
		return Part{}, errNotEnoughPartElems
	}

	for i := range split {
		keyValue := strings.Split(split[i], "=")
		if len(keyValue) != 2 {
			return Part{}, errInvalidKeyValue
		}

		value, err := strconv.Atoi(keyValue[1])
		if err != nil {
			return Part{}, err
		}

		switch keyValue[0] {
		case "x":
			part.X = value
		case "m":
			part.M = value
		case "a":
			part.A = value
		case "s":
			part.S = value
		}
	}

	return part, nil
}
