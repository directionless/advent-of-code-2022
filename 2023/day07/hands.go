package day07

import (
	"bytes"
	"fmt"
)

type Hand struct {
	Bid         int
	val         handValue
	comparable1 []byte
	comparable2 []byte
}

func HandFromBytes(in []byte, bid int) Hand {
	if len(in) != 5 {
		panic("This is not a hand. Does not have 5 cards")
	}

	val := valueFromHand(in, false)

	return Hand{
		Bid: bid,
		val: val,
		comparable1: []byte{
			byte(val),
			byteToComparable(in[0], false),
			byteToComparable(in[1], false),
			byteToComparable(in[2], false),
			byteToComparable(in[3], false),
			byteToComparable(in[4], false),
		},
		comparable2: []byte{
			byte(valueFromHand(in, true)),
			byteToComparable(in[0], true),
			byteToComparable(in[1], true),
			byteToComparable(in[2], true),
			byteToComparable(in[3], true),
			byteToComparable(in[4], true),
		},
	}
}

func (h Hand) String() string {
	return fmt.Sprintf("%s", []byte{
		comparableToByte(h.comparable1[1]),
		comparableToByte(h.comparable1[2]),
		comparableToByte(h.comparable1[3]),
		comparableToByte(h.comparable1[4]),
		comparableToByte(h.comparable1[5]),
	})
}

func (h Hand) LexComparable1() []byte {
	return h.comparable1
}

func (h Hand) LexComparable2() []byte {
	return h.comparable2
}

func byteToComparable(in byte, part2 bool) byte {
	switch {
	case byte('2') <= in && in <= byte('9'):
		return in
	case in == byte('T'):
		return byte('a')
	case in == byte('J'):
		if part2 {
			return byte('1')
		}
		return byte('b')
	case in == byte('Q'):
		return byte('c')
	case in == byte('K'):
		return byte('d')
	case in == byte('A'):
		return byte('e')
	}

	panic("This is not an known card -- byteToComparable")
}

func comparableToByte(in byte) byte {
	switch {
	case byte('2') <= in && in <= byte('9'):
		return in
	case in == byte('a'):
		return byte('T')
	case in == byte('b') || in == byte('1'):
		return byte('J')
	case in == byte('c'):
		return byte('Q')
	case in == byte('d'):
		return byte('K')
	case in == byte('e'):
		return byte('A')
	}

	panic("This is not an known card -- comparableToByte")
}

type handValue byte

const (
	highCard     handValue = handValue('a')
	onePair      handValue = handValue('b')
	twoPair      handValue = handValue('c')
	threeOfAKind handValue = handValue('d')
	fullHouse    handValue = handValue('e')
	fourOfAKind  handValue = handValue('f')
	fiveOfAKind  handValue = handValue('g')
)

var jokerMap = map[handValue][5]handValue{
	highCard:     {highCard, onePair, threeOfAKind, fourOfAKind, fiveOfAKind},
	onePair:      {onePair, threeOfAKind, fourOfAKind, fiveOfAKind, fiveOfAKind},
	twoPair:      {twoPair, fullHouse, fourOfAKind, fiveOfAKind, fiveOfAKind},
	threeOfAKind: {threeOfAKind, fourOfAKind, fiveOfAKind, fiveOfAKind, fiveOfAKind},
	fullHouse:    {fullHouse, fourOfAKind, fiveOfAKind, fiveOfAKind, fiveOfAKind},
	fourOfAKind:  {fourOfAKind, fiveOfAKind, fiveOfAKind, fiveOfAKind, fiveOfAKind},
	fiveOfAKind:  {fiveOfAKind, fiveOfAKind, fiveOfAKind, fiveOfAKind, fiveOfAKind},
}

func valueFromHand(in []byte, withJokers bool) handValue {
	// Hand value is based on power hands. pair, 3 of a kind, etc.
	// So the first thing we need to do is look for pair, triples, etc.
	// Their value doen't matter, we just need to know how many we have.
	jokers := 0
	counts := make(map[byte]int, 5)
	for i := 0; i < 5; i++ {
		if withJokers && in[i] == byte('J') {
			jokers += 1
			continue
		}
		counts[in[i]] += 1
	}

	// Then, we need to figure out which is which.
	// As there are only 7 kinds of hand, we can just create a manual routine.
	ret := highCard
	if len(counts) == 5 {
		ret = highCard
	}

	if len(counts) == 1 {
		ret = fiveOfAKind
	}

	pairs := 0
	triples := 0
	for _, c := range counts {
		switch c {
		case 2:
			pairs += 1
		case 3:
			triples += 1
		case 4:
			ret = fourOfAKind
		}
	}

	switch {
	case pairs == 1 && triples == 1:
		ret = fullHouse
	case pairs == 1:
		ret = onePair
	case pairs == 2:
		ret = twoPair
	case triples == 1:
		ret = threeOfAKind
	}

	//fmt.Printf("in: %s;  counts: %v; pairs: %d; triples: %d, ret: %s, jokers: %d\n", in, counts, pairs, triples, string(ret), jokers)

	if jokers == 5 {
		return fiveOfAKind
	}

	return jokerMap[ret][jokers]

	//fmt.Printf("in: %s;  counts: %v; pairs: %d; triples: %d\n", in, counts, pairs, triples)
	//panic("unable to calculate hand value")
}

type handSorter1 []Hand

// Len is part of sort.Interface.
func (s handSorter1) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s handSorter1) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s handSorter1) Less(i, j int) bool {
	return bytes.Compare(s[i].LexComparable1(), s[j].LexComparable1()) < 0
}

type handSorter2 []Hand

// Len is part of sort.Interface.
func (s handSorter2) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s handSorter2) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s handSorter2) Less(i, j int) bool {
	return bytes.Compare(s[i].LexComparable2(), s[j].LexComparable2()) < 0
}
