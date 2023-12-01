package day17

import (
	"fmt"
)


type dayHandler struct {
	jetPattern []byte
	chamberTop [chamberWidth]int
}

func New() *dayHandler {
	h := &dayHandler{}

	return h

}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	for _, c := range line {
		if c != '<' && c != '>' {
			return fmt.Errorf("invalid character %c", c)
		}
	}

	h.jetPattern = line

	return nil
}

func

func (h *dayHandler) AnswerPart1() int {



	return 0
}

func (h *dayHandler) AnswerPart2() int {
	return 0
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
