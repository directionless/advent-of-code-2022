package day9

import (
	"errors"
	"fmt"
	"strconv"
)

// point is an xy point. I needed it because go was really unhappy with map[int][2]int,
// and map[int]*point was simple
type point struct {
	Row int
	Col int
}

func (p point) Array() [2]int {
	return [2]int{p.Row, p.Col}
}

type dayHandler struct {
	len    int
	visted map[int]map[[2]int]bool
	pos    map[int]*point

	vm *videoMaker
}

func New() *dayHandler {
	length := 10
	h := &dayHandler{
		len:    length,
		visted: make(map[int]map[[2]int]bool, length),
		pos:    make(map[int]*point, length),
	}

	for l := 0; l <= length; l++ {
		h.pos[l] = &point{}
		h.visted[l] = make(map[[2]int]bool)
		h.visted[l][h.pos[l].Array()] = true
	}

	if h.vm != nil {
		h.vm.AddMove(h.pos)
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

	// distance might be multicharacters, and as these are string encoded
	// anyhow, Atoi(string()) is just simpler.
	totalDistance, err := strconv.Atoi(string(line[2:]))
	if err != nil {
		return fmt.Errorf("unable to convert to distance: %w", err)
	}

	for d := 0; d < totalDistance; d++ {
		fmt.Printf("\nline %s 1: headPos: %v, seg1: %v\n", line, h.pos[0], h.pos[1])

		// Move head _only_ 1 at a time.
		switch line[0] {
		case byte('U'):
			h.pos[0].Col += 1
		case byte('D'):
			h.pos[0].Col -= 1
		case byte('L'):
			h.pos[0].Row -= 1
		case byte('R'):
			h.pos[0].Row += 1
		default:
			return errors.New("unknown direction")
		}

		//fmt.Printf("line %s 2: headPos: %v, tailPos: %v\n", line, h.headPos, h.tailPos)

		// With the head moved, now move each segment to follow.
		for seg := 1; seg < h.len; seg++ {
			if err := h.moveSegment(seg); err != nil {
				return fmt.Errorf("moving seg %d: %w", seg, err)
			}
		}

		//fmt.Printf("line %s 3: headPos: %v, seg1: %v\n", line, h.pos[0], h.pos[1])

	}

	//fmt.Printf("after line %s: headPos: %v, seg1: %v\n", line, h.pos[0], h.pos[1])
	return nil
}

// moveSegment moves a segment to be closer to the preceeding segment,
// according to the provided rules.
func (h *dayHandler) moveSegment(seg int) error {
	h.visted[seg][h.pos[seg].Array()] = true

	ropeLen := 1

	//fmt.Printf("moveTail: rowDiff: %d, colDiff: %d\n", h.headPos[0]-h.tailPos[0], h.headPos[1]-h.tailPos[1])

	// Movement is defined as:
	//
	// If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move one step
	// in that direction so it remains close enough:
	//
	// Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves
	// one step diagonally to keep up:

	rowDiff := h.pos[seg-1].Row - h.pos[seg].Row
	colDiff := h.pos[seg-1].Col - h.pos[seg].Col

	// Early return if there is any touching
	if abs(rowDiff) <= ropeLen && abs(colDiff) <= ropeLen {
		return nil
	}

	// Rows
	switch {
	case rowDiff == 0:
		// nothing
	case colDiff == 0 && rowDiff > ropeLen:
		h.pos[seg].Row += 1
	case colDiff == 0 && rowDiff < -ropeLen:
		h.pos[seg].Row -= 1
	case rowDiff >= ropeLen:
		h.pos[seg].Row += 1
	case rowDiff <= ropeLen:
		h.pos[seg].Row -= 1
	}

	// move column
	switch {
	case colDiff == 0:
		// nothing
	case rowDiff == 0 && colDiff > 1:
		h.pos[seg].Col += 1
	case rowDiff == 0 && colDiff < 1:
		h.pos[seg].Col -= 1
	case colDiff >= 1:
		h.pos[seg].Col += 1
	case colDiff <= 1:
		h.pos[seg].Col -= 1
	}

	h.visted[seg][h.pos[seg].Array()] = true

	return nil
}

func (h *dayHandler) AnswerPart1() int {
	//fmt.Printf("Head is now row,col: %v\n", h.headPos)
	//fmt.Printf("Tail is now row,col: %v\n", h.tailPos)

	//fmt.Printf("%v\n", h.tailVisited)

	// Make sure we denote that we visited the starting point
	//	h.visited[1]ited[[2]int{0, 0}] = true

	return len(h.visted[1])
}

func (h *dayHandler) AnswerPart2() int {
	return len(h.visted[9])
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: segment1 visted: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: segment9 visted: %d\n", h.AnswerPart2())
}

// abs is an absolute value function for ints.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
