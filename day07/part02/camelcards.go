package part02

import (
	"bufio"
	"cmp"
	"slices"
	"strconv"
	"strings"
)

const minAlloc = 64

type Hand struct {
	Bid   int
	Cards [5]byte
}

func Parse(input string) []Hand {
	if input == "" {
		return nil
	}

	scanner := bufio.NewScanner(strings.NewReader(input))

	hands := make([]Hand, 0, minAlloc)

	for scanner.Scan() {
		if hand, ok := parse(scanner.Text()); ok {
			hands = append(hands, hand)
		}
	}

	return hands
}

func Rank(hands ...Hand) int {
	kinds := make(map[int][]Hand, len(hands))
	kindValues := make([]int, 0, len(hands))

	for i := range hands {
		kindValue := kind(hands[i].Cards)

		if _, ok := kinds[kindValue]; !ok {
			kinds[kindValue] = make([]Hand, 0, minAlloc)
			kindValues = append(kindValues, kindValue)
		}

		kinds[kindValue] = append(kinds[kindValue], hands[i])
	}

	slices.Sort(kindValues)

	var (
		result int
		rank   = 1
	)

	for i := 0; i < len(kindValues); i++ {
		switch len(kinds[kindValues[i]]) {
		case 1:
			result += kinds[kindValues[i]][0].Bid * rank
			rank++
		default:
			slices.SortFunc(kinds[kindValues[i]], func(a, b Hand) int {
				var comp int

				for idx := range a.Cards {
					comp = cmp.Compare(value(a.Cards[idx]), value(b.Cards[idx]))
					if comp != 0 {
						break
					}
				}

				return comp
			})

			for idx := 0; idx < len(kinds[kindValues[i]]); idx++ {
				result += kinds[kindValues[i]][idx].Bid * rank
				rank++
			}
		}
	}

	return result
}

func parse(line string) (Hand, bool) {
	if line == "" {
		return Hand{}, false
	}

	fields := strings.Fields(line)

	if len(fields) != 2 || len(fields[0]) != 5 {
		return Hand{}, false
	}

	bid, err := strconv.Atoi(fields[1])
	if err != nil {
		return Hand{}, false
	}

	return Hand{
		Bid:   bid,
		Cards: [5]byte([]byte(fields[0])),
	}, true
}

const (
	highCard     = 1
	onePair      = 2
	twoPair      = 3
	threeOfAKind = 4
	fullHouse    = 5
	fourOfAKind  = 6
	fiveOfAKind  = 7
)

func kind(hand [5]byte) int {
	m := make(map[byte]int, len(hand))

	for i := range hand {
		m[hand[i]]++
	}

	combined := make([]int, 0, len(m))

	var jokers int
	for k, v := range m {
		if k == 'J' {
			jokers = v

			continue
		}

		combined = append(combined, v)
	}

	slices.Sort(combined)

	switch jokers {
	case 5:
		combined = append(combined, jokers)
	case 0:
		break
	default:
		combined[len(combined)-1] += jokers
	}

	switch {
	case len(combined) == 1 && combined[len(combined)-1] == 5:
		return fiveOfAKind
	case len(combined) == 2 && combined[len(combined)-1] == 4:
		return fourOfAKind
	case len(combined) == 2 && combined[len(combined)-1] == 3:
		return fullHouse
	case len(combined) == 3 && combined[len(combined)-1] == 3:
		return threeOfAKind
	case len(combined) == 3 && combined[len(combined)-1] == 2 && combined[len(combined)-2] == 2:
		return twoPair
	default:
		for i := range combined {
			if combined[i] == 2 {
				return onePair
			}
		}

		return highCard
	}
}

func value(card byte) int {
	switch card {
	case 'J':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'Q':
		return 11
	case 'K':
		return 12
	case 'A':
		return 13
	default:
		return 0
	}
}
