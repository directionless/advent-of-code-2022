package day01

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const ()

var (
	findNumbers     = regexp.MustCompile(`(\d)`)
	findNumberWords = regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)
)

func wordToNum(w []byte) []byte {
	switch {
	case bytes.Compare([]byte("one"), w) == 0:
		return []byte("1")
	case bytes.Compare([]byte("two"), w) == 0:
		return []byte("2")
	case bytes.Compare([]byte("three"), w) == 0:
		return []byte("3")
	case bytes.Compare([]byte("four"), w) == 0:
		return []byte("4")
	case bytes.Compare([]byte("five"), w) == 0:
		return []byte("5")
	case bytes.Compare([]byte("six"), w) == 0:
		return []byte("6")
	case bytes.Compare([]byte("seven"), w) == 0:
		return []byte("7")
	case bytes.Compare([]byte("eight"), w) == 0:
		return []byte("8")
	case bytes.Compare([]byte("nine"), w) == 0:
		return []byte("9")
	}
	return nil
}

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
	if i, err := h.lineToNum(line); err == nil {
		h.running_total += i
	}

	// part 2
	linep2 := findNumberWords.ReplaceAllFunc(line, wordToNum)
	i2, err := h.lineToNum(linep2)
	if err != nil {
		return fmt.Errorf("parsing p2 number from %s: %w", linep2, err)
	}

	fmt.Println(string(line), string(linep2), i2)

	h.running_total_p2 += i2
	return nil
}

func (h *dayHandler) lineToNum(line []byte) (int, error) {
	m := findNumbers.FindAll(line, -1)
	if len(m) < 1 {
		return 0, errors.New("no numbers found")
	}

	num := fmt.Sprintf("%s%s", m[0], m[len(m)-1])

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
