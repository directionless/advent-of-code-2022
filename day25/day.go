package day25

import (
	"fmt"

	"github.com/directionless/advent-of-code-2022/day25/snafu"
)

const ()

type dayHandler struct {
	sum int
}

func New() *dayHandler {
	h := &dayHandler{}

	return h

}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	num, err := snafu.ToInt(string(line))
	if err != nil {
		return fmt.Errorf("unable to parse %s: %w", line, err)
	}

	h.sum += num

	return nil
}

func (h *dayHandler) AnswerPart1() int {
	return h.sum
}

func (h *dayHandler) AnswerPart1Snafu() (string, error) {
	s, err := snafu.FromInt(h.sum)
	if err != nil {
		return "", err
	}

	return s, err
}

func (h *dayHandler) AnswerPart2() int {
	return 0
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
