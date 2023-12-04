package part01

import (
	"strconv"
	"strings"
)

type Scratchcard struct {
	id      int
	winning map[int]struct{}
	input   []int
}

func (s Scratchcard) Matches() int {
	var n int

	for i := range s.input {
		if _, ok := s.winning[s.input[i]]; ok {
			n++
		}
	}

	return n
}

func Sum(cards ...Scratchcard) int {
	var n int

	for i := range cards {
		var (
			m       = 1
			matches = cards[i].Matches()
		)

		switch matches {
		case 0:
			continue
		case 1:
			n++
		default:
			// subtract 1 for m = 1
			matches--

			for ; matches > 0; matches-- {
				m *= 2
			}

			n += m
		}
	}

	return n
}

func NewSet(input string) []Scratchcard {
	lines := strings.Split(input, "\n")

	cards := make([]Scratchcard, 0, len(lines))

	for i := range lines {
		var card Scratchcard

		idSplit := strings.Split(lines[i], ":")
		if len(idSplit) != 2 {
			continue
		}

		if idFields := strings.Fields(idSplit[0]); len(idFields) == 2 {
			// error is less important since the ID doesn't matter
			card.id, _ = strconv.Atoi(idFields[1])
		}

		cardsSplit := strings.Split(idSplit[1], "|")
		if len(cardsSplit) != 2 {
			continue
		}

		winningCards := strings.Fields(cardsSplit[0])
		card.winning = make(map[int]struct{}, len(winningCards))

		for idx := range winningCards {
			if value, err := strconv.Atoi(winningCards[idx]); err == nil {
				card.winning[value] = struct{}{}
			}
		}

		playingCards := strings.Fields(cardsSplit[1])
		card.input = make([]int, 0, len(playingCards))

		for idx := range playingCards {
			if value, err := strconv.Atoi(playingCards[idx]); err == nil {
				card.input = append(card.input, value)
			}
		}

		cards = append(cards, card)
	}

	return cards
}
