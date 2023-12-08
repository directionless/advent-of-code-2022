package day08

import (
	"fmt"
)

const (
	ExampleAnswer1 = -1
	ExampleAnswer2 = -1

	RealAnswer1 = -1
	RealAnswer2 = -1
)

type dayHandler struct {
	part1Answer any
	part2Answer any
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

// Solve is called when the input is done being Consumed. Some puzzle can be solved entirely
// in Consume, line by line. Others need an additional step
func (h *dayHandler) Solve() error {
	// This has the handy effect of catchig the EOF. Advance lines, and call solve.
	return nil
}

func (h *dayHandler) AnswerPart1() any {
	return h.part1Answer

}

func (h *dayHandler) AnswerPart2() any {
	return h.part2Answer

}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
