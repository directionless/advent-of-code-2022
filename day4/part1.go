package day4

import (
	"bytes"
	"fmt"
	"strconv"
)

type Part1Handler struct {
	countOfConsumedPairs int
	partiallyOverlapping int
}

func NewPart1() *Part1Handler {
	return &Part1Handler{}
}

func (h *Part1Handler) Consume(line []byte) error {
	//fmt.Printf("line: %s\n", string(line))

	pairs := bytes.Split(line, []byte(","))
	if len(pairs) != 2 {
		return fmt.Errorf("invalid line: %s", line)
	}

	min1, max1, err := bytesToRange(pairs[0])
	if err != nil {
		return fmt.Errorf("invalid range: %s", string(pairs[0]))
	}

	min2, max2, err := bytesToRange(pairs[1])
	if err != nil {
		return fmt.Errorf("invalid range: %s", string(pairs[0]))
	}

	// fmt.Printf("%d - %d :: %d - %d\n", min1, max1, min2, max2)

	if isSubSet(min1, max1, min2, max2) || isSubSet(min2, max2, min1, max1) {
		h.countOfConsumedPairs++
	}

	if anyOverlap(min1, max1, min2, max2) || anyOverlap(min2, max2, min1, max1) {
		h.partiallyOverlapping++
	}

	return nil
}

func (h *Part1Handler) Answer() int {
	return h.countOfConsumedPairs
}

func (h *Part1Handler) AnswerPart2() int {
	return h.partiallyOverlapping
}

func (h *Part1Handler) Print() {
	fmt.Printf("Elves with fully contained pairs: %d\n", h.Answer())
}

func isSubSet(min1, max1, min2, max2 int) bool {
	// I wrote this weird so I can add some printf debugs
	if min1 > min2 {
		//fmt.Printf("min1 > min2\n")
		return false
	}

	if max1 < max2 {
		//fmt.Printf("max1 < max2\n")
		return false
	}

	return true
}

func anyOverlap(min1, max1, min2, max2 int) bool {
	if min1 >= min2 && min1 <= max2 {
		return true
	}

	if max1 >= min2 && max1 <= max2 {
		return true
	}

	return false
}

func bytesToRange(b []byte) (int, int, error) {
	parts := bytes.Split(b, []byte("-"))
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid range: %s. Got: %v", b, parts)
	}

	min, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert %s to int", string(parts[0]))
	}

	max, err := strconv.Atoi(string(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert %s to int", string(parts[1]))
	}

	return min, max, nil
}
