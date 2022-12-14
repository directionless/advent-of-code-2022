package day13

import (
	"fmt"
	"sort"

	"github.com/davecgh/go-spew/spew"
)

const ()

type dayHandler struct {
	left       listNumber
	right      listNumber
	pairIndex  int
	part1Sum   int
	allPackets []listNumber
	divider1   listNumber
	divider2   listNumber
}

func New() *dayHandler {
	h := &dayHandler{}

	var err error
	h.divider1, err = ParseNumber([]byte("[[2]]"))
	if err != nil {
		panic(err)
	}
	h.divider2, err = ParseNumber([]byte("[[6]]"))
	if err != nil {
		panic(err)
	}

	return h

}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	switch {
	case h.left == nil:
		parsed, err := ParseNumber(line)
		if err != nil {
			return fmt.Errorf("could not parse number %s: %w", line, err)
		}
		h.left = parsed
		h.allPackets = append(h.allPackets, parsed)
	case h.right == nil:
		h.pairIndex++

		parsed, err := ParseNumber(line)
		if err != nil {
			return fmt.Errorf("could not parse number %s: %w", line, err)
		}
		h.right = parsed
		h.allPackets = append(h.allPackets, parsed)

		ret, err := CompareNumbers(h.left, h.right)
		if err != nil {
			return fmt.Errorf("could not compare numbers: %w", err)
		}

		// What are the indices of the pairs that are already in the right order? (The first pair has index 1, the second pair has index 2, and so on.) In the above example, the pairs in the right order are 1, 2, 4, and 6; the sum of these indices is 13.
		if ret == -1 {
			h.part1Sum += h.pairIndex
		}

		// clear
		h.left = nil
		h.right = nil

	case len(line) == 0:
		h.left = nil
		h.right = nil
	}

	return nil
}

func (h *dayHandler) AnswerPart1() int {
	return h.part1Sum
}

func (h *dayHandler) AnswerPart2() int {
	sorter := func(i, j int) bool {
		ret, err := CompareNumbers(h.allPackets[i], h.allPackets[j])
		if err != nil {
			panic(err)
		}
		return ret == -1
	}

	sort.Slice(h.allPackets, sorter)

	for i, num := range h.allPackets {
		if comp, _ := CompareNumbers(num, h.divider1); comp == 0 {
			fmt.Printf("found divider1: %d\n", i)
			continue
		}
		if comp, _ := CompareNumbers(num, h.divider2); comp == 0 {
			fmt.Printf("found divider2: %d\n", i)
			continue
		}
	}

	spew.Dump(h.allPackets[11])
	return 0
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
