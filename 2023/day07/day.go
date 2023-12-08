package day07

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

const (
	ExampleAnswer1 = 6440
	ExampleAnswer2 = 5905

	RealAnswer1 = 253954294
	RealAnswer2 = 254837398 // 254141271 too low
)

type dayHandler struct {
	hands []Hand
}

func New() *dayHandler {
	h := &dayHandler{}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	lineSplit := bytes.SplitN(line, []byte(" "), 2)
	if len(lineSplit) != 2 {
		return fmt.Errorf(`line %s had %d chunks, expected two`, line, len(lineSplit))
	}

	bid, err := strconv.Atoi(string(lineSplit[1]))
	if err != nil {
		return fmt.Errorf(`line "%s" convertin to int: %w`, line, err)
	}

	h.hands = append(h.hands, HandFromBytes(lineSplit[0], bid))

	return nil
}

// Solve is called when the input is done being Consumed. Some puzzle can be solved entirely
// in Consume, line by line. Others need an additional step
func (h *dayHandler) Solve() error {

	return nil

}

func (h *dayHandler) AnswerPart1() any {
	sort.Sort(handSorter1(h.hands))

	//for _, hand := range h.hands {
	//	fmt.Printf("%s: %s\n", hand.String(), hand.LexComparable())
	//}

	// Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.
	winnings := 0
	for i, hand := range h.hands {
		rank := i + 1
		winnings += rank * hand.Bid

	}

	return winnings

}

func (h *dayHandler) AnswerPart2() any {
	sort.Sort(handSorter2(h.hands))

	//for _, hand := range h.hands {
	//	fmt.Printf("%s: %s\n", hand.String(), hand.LexComparable2())
	//}

	winnings := 0
	for i, hand := range h.hands {
		rank := i + 1
		winnings += rank * hand.Bid

	}

	return winnings

}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
