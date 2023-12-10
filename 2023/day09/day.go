package day09

import (
	"fmt"
)

const (
	ExampleAnswer1 = 114
	ExampleAnswer2 = 2

	RealAnswer1 = 1772145754
	RealAnswer2 = 867
)

type dayHandler struct {
	sequences []*sequence
	reversed  []*sequence

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
	{
		s, err := sequenceFromLine(line)
		if err != nil {
			return fmt.Errorf(`parsing "%s": %w`, line, err)
		}

		if err := s.Solve(); err != nil {
			return err
		}

		h.sequences = append(h.sequences, s)
	}

	{
		s, err := sequenceFromLine(line)
		if err != nil {
			return fmt.Errorf(`parsing "%s": %w`, line, err)
		}

		s.Reverse()

		if err := s.Solve(); err != nil {
			return err
		}

		h.reversed = append(h.reversed, s)

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
	tot := 0
	for _, s := range h.sequences {
		n, err := s.FindNext(1)
		if err != nil {
			panic(err)
		}
		tot += n
	}

	return tot
}

func (h *dayHandler) AnswerPart2() any {
	tot := 0
	for _, s := range h.reversed {
		n, err := s.FindNext(1)
		if err != nil {
			panic(err)
		}
		tot += n
	}

	return tot

}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
