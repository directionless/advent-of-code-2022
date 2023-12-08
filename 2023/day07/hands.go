package day07

import (
	"bytes"
	"fmt"
)

type Hand struct {
	Bid        int
	val        handValue
	comparable []byte
}

func HandFromBytes(in []byte, bid int) Hand {
	if len(in) != 5 {
		panic("This is not a hand. Does not have 5 cards")
	}

	val := valueFromHand(in)

	return Hand{
		Bid: bid,
		val: val,
		comparable: []byte{
			byte(val),
			byteToComparable(in[0]),
			byteToComparable(in[1]),
			byteToComparable(in[2]),
			byteToComparable(in[3]),
			byteToComparable(in[4]),
		},
	}
}

func (h Hand) String() string {
	return fmt.Sprintf("%s", []byte{
		comparableToByte(h.comparable[1]),
		comparableToByte(h.comparable[2]),
		comparableToByte(h.comparable[3]),
		comparableToByte(h.comparable[4]),
		comparableToByte(h.comparable[5]),
	})
}

func (h Hand) LexComparable() []byte {
	return h.comparable
}

func byteToComparable(in byte) byte {
	switch {
	case byte('2') <= in && in <= byte('9'):
		return in
	case in == byte('T'):
		return byte('a')
	case in == byte('J'):
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
	case in == byte('b'):
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

// Compare compares hands
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func CompareHands(a, b Hand) int {
	return bytes.Compare(a.LexComparable(), b.LexComparable())
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

func valueFromHand(in []byte) handValue {
	// Hand value is based on power hands. pair, 3 of a kind, etc.
	// So the first thing we need to do is look for pair, triples, etc.
	// Their value doen't matter, we just need to know how many we have.
	counts := make(map[byte]int, 5)
	for i := 0; i < 5; i++ {
		counts[in[i]] += 1
	}

	// Then, we need to figure out which is which.
	// As there are only 7 kinds of hand, we can just create a manual routine.

	if len(counts) == 5 {
		return highCard
	}

	if len(counts) == 1 {
		return fiveOfAKind
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
			return fourOfAKind
		}
	}

	switch {
	case pairs == 1 && triples == 1:
		return fullHouse
	case pairs == 1:
		return onePair
	case pairs == 2:
		return twoPair
	case triples == 1:
		return threeOfAKind
	}

	fmt.Printf("in: %s;  counts: %v; pairs: %d; triples: %d\n", in, counts, pairs, triples)
	panic("unable to calculate hand value")
}

type handSorter []Hand

// Len is part of sort.Interface.
func (s handSorter) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s handSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s handSorter) Less(i, j int) bool {
	return CompareHands(s[i], s[j]) < 0
}
