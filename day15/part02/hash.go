package part02

import (
	"strconv"
	"strings"
)

const minAlloc = 64

type Op struct {
	hash  int
	key   string
	op    int
	value int
}

func Parse(input string) []Op {
	if input == "" {
		return nil
	}

	input = strings.ReplaceAll(input, "\n", "")
	raw := strings.Split(input, ",")

	ops := make([]Op, 0, len(raw))

	for i := range raw {
		if raw[i] == "" {
			continue
		}

		ops = append(ops, extract(raw[i]))
	}

	return ops
}

func extract(input string) Op {
	key := make([]byte, 0, len(input))
	var value int
	var i int

	for ; i < len(input); i++ {
		if input[i] == '=' || input[i] == '-' {
			break
		}

		key = append(key, input[i])
	}

	if i+1 < len(input) {
		v, err := strconv.Atoi(input[i+1:])

		if err == nil {
			value = v
		}
	}

	return Op{
		key:   string(key),
		hash:  Hash(string(key)),
		op:    int(input[i]),
		value: value,
	}
}

func Map(input ...Op) map[int][]Op {
	m := make(map[int][]Op)

	for i := range input {
		switch input[i].op {
		case '-':
			if slice, ok := m[input[i].hash]; ok {
				for idx := range slice {
					if slice[idx].key == input[i].key {
						switch {
						case idx+1 <= len(slice):
							slice = append(slice[:idx], slice[idx+1:]...)
						default:
							slice = slice[:idx]
						}

						break
					}
				}

				m[input[i].hash] = slice
			}
		case '=':
			slice, ok := m[input[i].hash]
			if !ok {
				m[input[i].hash] = make([]Op, 0, minAlloc)
				m[input[i].hash] = append(m[input[i].hash], input[i])

				continue
			}

			var changed bool

			for idx := range slice {
				if slice[idx].key == input[i].key {
					slice[idx] = input[i]

					changed = true

					break
				}
			}

			if !changed {
				slice = append(slice, input[i])
			}

			m[input[i].hash] = slice
		}
	}

	return m
}

func Sum(m map[int][]Op) int {
	var n int

	for k, v := range m {
		for i := range v {
			n += (k + 1) * (i + 1) * v[i].value
		}
	}

	return n
}

func Hash(input string) int {
	var n int

	for i := range input {
		n = hash(n + int(input[i]))
	}

	return n
}

func hash(value int) int {
	value *= 17
	return value % 256
}
