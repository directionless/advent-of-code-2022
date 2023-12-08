package day07

import (
	"fmt"
)

type comparableHand [5]byte

type Hand struct {
	val        handValue
	comparable comparableHand
}

func HandFromBytes(in []byte) Hand {
	if len(in) != 5 {
		panic("This is not a hand. Does not have 5 cards")
	}

	return Hand{
		val: valueFromHand(in),
		comparable: [5]byte{
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
		comparableToByte(h.comparable[0]),
		comparableToByte(h.comparable[1]),
		comparableToByte(h.comparable[2]),
		comparableToByte(h.comparable[3]),
		comparableToByte(h.comparable[4]),
	})
}

func byteToComparable(in byte) byte {
	switch {
	case byte('2') <= in && in <= byte('9'):
		return in
	case in == byte('T'):
		return byte('9') + 1
	case in == byte('J'):
		return byte('9') + 2
	case in == byte('Q'):
		return byte('9') + 3
	case in == byte('K'):
		return byte('9') + 4
	case in == byte('A'):
		return byte('9') + 5
	}

	panic("This is not an known card")
}

func comparableToByte(in byte) byte {
	switch {
	case byte('2') <= in && in <= byte('9'):
		return in
	case in == byte('9')+1:
		return byte('T')
	case in == byte('9')+2:
		return byte('J')
	case in == byte('9')+3:
		return byte('Q')
	case in == byte('9')+4:
		return byte('K')
	case in == byte('9')+5:
		return byte('A')
	}

	panic("This is not an known card")
}

type handValue int

const (
	highCard     handValue = 0
	onePair      handValue = 1
	twoPair      handValue = 2
	threeOfAKind handValue = 3
	fullHouse    handValue = 4
	fourOfAKind  handValue = 5
	fiveOfAKind  handValue = 6
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
