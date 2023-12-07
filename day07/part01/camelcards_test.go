package part01

import (
	"testing"

	"github.com/zalgonoise/advent-2023/day07"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants []Hand
	}{
		{
			name: "Example",
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`,
			wants: []Hand{
				{Bid: 765, Cards: [5]byte{'3', '2', 'T', '3', 'K'}},
				{Bid: 684, Cards: [5]byte{'T', '5', '5', 'J', '5'}},
				{Bid: 28, Cards: [5]byte{'K', 'K', '6', '7', '7'}},
				{Bid: 220, Cards: [5]byte{'K', 'T', 'J', 'J', 'T'}},
				{Bid: 483, Cards: [5]byte{'Q', 'Q', 'Q', 'J', 'A'}},
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			hands := Parse(testcase.input)

			isEqual(t, len(testcase.wants), len(hands))

			for i := range testcase.wants {
				isEqual(t, testcase.wants[i].Bid, hands[i].Bid)
				isEqual(t, testcase.wants[i].Cards, hands[i].Cards)
			}
		})
	}
}

func TestValue(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input Hand
		wants int
	}{
		{
			name:  "Input/1",
			input: Hand{Bid: 765, Cards: [5]byte{'3', '2', 'T', '3', 'K'}},
			wants: onePair,
		},
		{
			name:  "Input/2",
			input: Hand{Bid: 684, Cards: [5]byte{'T', '5', '5', 'J', '5'}},
			wants: threeOfAKind,
		},
		{
			name:  "Input/3",
			input: Hand{Bid: 28, Cards: [5]byte{'K', 'K', '6', '7', '7'}},
			wants: twoPair,
		},
		{
			name:  "Input/4",
			input: Hand{Bid: 220, Cards: [5]byte{'K', 'T', 'J', 'J', 'T'}},
			wants: twoPair,
		},
		{
			name:  "Input/5",
			input: Hand{Bid: 483, Cards: [5]byte{'Q', 'Q', 'Q', 'J', 'A'}},
			wants: threeOfAKind,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			v := kind(testcase.input.Cards)

			isEqual(t, testcase.wants, v)
		})
	}
}

func TestRank(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Input",
			input: day07.Input,
			wants: 250474325,
		},
		{
			name: "Example",
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`,
			wants: 6440,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			hands := Parse(testcase.input)

			isEqual(t, testcase.wants, Rank(hands...))
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
