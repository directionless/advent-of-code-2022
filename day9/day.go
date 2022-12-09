package day9

import (
	"errors"
	"fmt"
	"strconv"
)

const ()

type dayHandler struct {
	tailVisited map[[2]int]bool

	headPos [2]int
	tailPos [2]int
}

func New() *dayHandler {
	h := &dayHandler{
		tailVisited: make(map[[2]int]bool),
	}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	if len(line) < 3 {
		return fmt.Errorf("line too short: %s", line)
	}

	totalDistance, err := strconv.Atoi(string(line[2:]))
	if err != nil {
		return fmt.Errorf("unable to convert to distance: %w", err)
	}

	for d := 0; d < totalDistance; d++ {
		fmt.Printf("\nline %s 1: headPos: %v, tailPos: %v\n", line, h.headPos, h.tailPos)

		// Move head _only_ 1 at a time.
		switch line[0] {
		case byte('U'):
			h.headPos[1] += 1
		case byte('D'):
			h.headPos[1] -= 1
		case byte('L'):
			h.headPos[0] -= 1
		case byte('R'):
			h.headPos[0] += 1
		default:
			return errors.New("unknown direction")
		}

		//fmt.Printf("line %s 2: headPos: %v, tailPos: %v\n", line, h.headPos, h.tailPos)

		if err := h.moveTail(); err != nil {
			return fmt.Errorf("moving tail: %w", err)
		}

		fmt.Printf("line %s 3: headPos: %v, tailPos: %v\n", line, h.headPos, h.tailPos)

	}

	fmt.Printf("after line %s: headPos: %v, tailPos: %v\n", line, h.headPos, h.tailPos)

	return nil
}

// moveTail moves the tail according to the provided rules:
//
// Due to the aforementioned Planck lengths, the rope must be quite short; in fact, the head (H) and tail (T) must
// always be touching (diagonally adjacent and even overlapping both count as touching):
func (h *dayHandler) moveTail() error {
	ropeLen := 1

	//fmt.Printf("moveTail: rowDiff: %d, colDiff: %d\n", h.headPos[0]-h.tailPos[0], h.headPos[1]-h.tailPos[1])

	// Movement is defined as:
	//
	// If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move one step
	// in that direction so it remains close enough:
	//
	// Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves
	// one step diagonally to keep up:

	rowDiff := h.headPos[0] - h.tailPos[0]
	colDiff := h.headPos[1] - h.tailPos[1]

	// Early return if there is any touching
	if abs(rowDiff) <= ropeLen && abs(colDiff) <= ropeLen {
		return nil
	}

	// Rows
	switch {
	case rowDiff == 0:
		// nothing
	case colDiff == 0 && rowDiff > ropeLen:
		h.tailPos[0] += 1
	case colDiff == 0 && rowDiff < -ropeLen:
		h.tailPos[0] -= 1
	case rowDiff >= ropeLen:
		h.tailPos[0] += 1
	case rowDiff <= ropeLen:
		h.tailPos[0] -= 1
	}

	// move column
	switch {
	case colDiff == 0:
		// nothing
	case rowDiff == 0 && colDiff > 1:
		h.tailPos[1] += 1
	case rowDiff == 0 && colDiff < 1:
		h.tailPos[1] -= 1

	case colDiff >= 1:
		h.tailPos[1] += 1
	case colDiff <= 1:
		h.tailPos[1] -= 1
	}

	h.tailVisited[h.tailPos] = true

	return nil
}

func (h *dayHandler) AnswerPart1() int {
	fmt.Printf("Head is now row,col: %v\n", h.headPos)
	fmt.Printf("Tail is now row,col: %v\n", h.tailPos)

	//fmt.Printf("%v\n", h.tailVisited)

	// Make sure we denote that we visited the starting point
	h.tailVisited[[2]int{0, 0}] = true

	return len(h.tailVisited)
}

func (h *dayHandler) AnswerPart2() int {
	return 0
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}

// abs is an absolute value function for ints.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
