package day12

import (
	"fmt"
)

const ()

type dayHandler struct {
	grid   *grid
	width  int
	height int
}

func New() *dayHandler {
	h := &dayHandler{
		grid: NewGrid(),
	}
	return h
}

func (h *dayHandler) Grid() *grid {
	return h.grid
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	h.grid.AddRow(line)

	return nil
}

func (h *dayHandler) AnswerPart1() int {
	h.grid.ResetBFS()
	return h.grid.BFS(false)
}

func (h *dayHandler) AnswerPart2() int {
	h.grid.ResetBFS()
	return h.grid.BFS(true)
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
