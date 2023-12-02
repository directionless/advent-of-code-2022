package day01

import (
	"fmt"
	"strconv"
)

type dayHandler struct {
	running_total    int
	running_total_p2 int
}

func New() *dayHandler {
	h := &dayHandler{}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	// On each line, the calibration value can be found by combining the first digit
	// and the last digit (in that order) to form a single two-digit number.

	// part1 isn't wholly real, so we ignore errors here. Messy, but :shrug:
	if i, err := h.lineToNum(string(line), false); err == nil {
		h.running_total += i
	}

	// part 2
	i2, err := h.lineToNum(string(line), true)
	if err != nil {
		return fmt.Errorf("parsing p2 number from %s: %w", line, err)
	}

	h.running_total_p2 += i2

	return nil
}

func (h *dayHandler) lineToNum(line string, withWords bool) (int, error) {
	var extras []map[string]any
	if withWords {
		extras = []map[string]any{numbersStr}
	}

	_, first := findFirst(line, numbersInt, extras...)
	_, last := findLast(line, numbersInt, extras...)

	if last == nil {
		last = first
	}

	num := fmt.Sprintf("%d%d", first, last)

	i, err := strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf("unable to convert %s to an int", num)
	}

	return i, nil
}

func (h *dayHandler) AnswerPart1() any {
	return h.running_total
}

func (h *dayHandler) AnswerPart2() any {
	return h.running_total_p2
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
