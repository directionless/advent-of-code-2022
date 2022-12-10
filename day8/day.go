package day8

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

const ()

type dayHandler struct {
	forrest      [][]byte
	visibleTrees map[[2]int]bool
	rowSize      int
	columnSize   int
}

func New() *dayHandler {
	h := &dayHandler{
		forrest:      [][]byte{},
		visibleTrees: make(map[[2]int]bool),
	}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	h.forrest = append(h.forrest, line)

	return nil
}

func (h *dayHandler) Check() error {
	h.rowSize = len(h.forrest[0])
	h.columnSize = len(h.forrest)

	for _, row := range h.forrest {
		if len(row) != h.rowSize {
			return fmt.Errorf("row size mismatch")
		}
	}
	return nil
}

func (h *dayHandler) Dump() {
	spew.Dump(h.forrest)
}

func (h *dayHandler) visibilityHorizontal() {
	var viewHeight byte

	for rowNum, row := range h.forrest {

		// visibility from the left
		viewHeight = 0 //row[0] // ignore first tree
		for i := 0; i < h.rowSize; i++ {
			if row[i] > viewHeight {
				treePos := [2]int{rowNum, i}
				h.visibleTrees[treePos] = true
				viewHeight = row[i]
			}
		}

		// visibility from the right
		viewHeight = 0 // row[h.rowSize-1] // ignore first tree
		for i := h.rowSize - 1; i >= 0; i-- {
			if row[i] > viewHeight {
				treePos := [2]int{rowNum, i}
				h.visibleTrees[treePos] = true
				viewHeight = row[i]
			}
		}

	}

}

func (h *dayHandler) visibilityVertical() {
	// This is harder to walk, since we need to slice things the opposite way

	var viewHeight byte

	for columnNum := 0; columnNum < h.rowSize; columnNum++ {

		// From the top
		viewHeight = 0 //h.forrest[0][columnNum] // ignore first tree
		for i := 0; i < len(h.forrest); i++ {
			if h.forrest[i][columnNum] > viewHeight {
				treePos := [2]int{i, columnNum}
				h.visibleTrees[treePos] = true
				viewHeight = h.forrest[i][columnNum]
			}
		}

		// From the bottom
		viewHeight = 0 //h.forrest[len(h.forrest)-1][columnNum] // ignore first tree
		for i := len(h.forrest) - 1; i >= 0; i-- {
			if h.forrest[i][columnNum] > viewHeight {
				treePos := [2]int{i, columnNum}
				h.visibleTrees[treePos] = true
				viewHeight = h.forrest[i][columnNum]
			}
		}
	}
}

type directionType string

const (
	left  directionType = "left"
	right directionType = "right"
	up    directionType = "up"
	down  directionType = "down"
)

// Look is used for scenic score.
func (h *dayHandler) Look(rowIdx, colIdx int, direction directionType) int {

	var looper = &looperStruct{Row: rowIdx, Col: colIdx}

	switch direction {
	case right:
		looper.End = h.rowSize - 1
		looper.step = "incr"
		looper.constRowOrCol = "row"
	case left:
		looper.End = 0
		looper.step = "decr"
		looper.constRowOrCol = "row"

	case down:
		looper.End = h.columnSize - 1
		looper.step = "incr"
		looper.constRowOrCol = "col"
	case up:
		looper.End = 0
		looper.step = "decr"
		looper.constRowOrCol = "col"
	}

	myTreeHeight := h.forrest[rowIdx][colIdx]
	var visibleTrees int

	for {
		looper.Next()

		if looper.Done() {
			break
		}

		targetTreeHeight := h.forrest[looper.Row][looper.Col]
		visibleTrees++

		if targetTreeHeight >= myTreeHeight {
			break
		}
	}

	return visibleTrees
}

func (h *dayHandler) ScenicScore(rowIdx, columnIdx int) int {
	return h.Look(rowIdx, columnIdx, right) * h.Look(rowIdx, columnIdx, left) * h.Look(rowIdx, columnIdx, up) * h.Look(rowIdx, columnIdx, down)
}

func (h *dayHandler) AnswerPart1() int {
	h.visibilityHorizontal()
	h.visibilityVertical()
	return len(h.visibleTrees)
}

func (h *dayHandler) AnswerPart2() int {
	mostScenic := 0
	atRow := 0
	atCol := 0

	for rowIdx := 0; rowIdx < h.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < h.columnSize; columnIdx++ {
			scenicScore := h.ScenicScore(rowIdx, columnIdx)

			if scenicScore > mostScenic {
				atRow = rowIdx
				atCol = columnIdx
				mostScenic = scenicScore
			}
		}
	}

	_ = atRow
	_ = atCol

	return mostScenic
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: The most scenic score %d\n", h.AnswerPart2())
}
