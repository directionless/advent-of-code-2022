package day04

import (
	"fmt"
)

type dayHandler struct {
}

func New() *dayHandler {
	h := &dayHandler{}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	return nil
}

func (h *dayHandler) Solve() error {

	return nil
}

func (h *dayHandler) AnswerPart1() any {
	return nil
}

func (h *dayHandler) AnswerPart2() any {
	return nil
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
