package day4

import "fmt"

type Part1Handler struct {
}

func NewPart1() *Part1Handler {
	return &Part1Handler{}
}

func (h *Part1Handler) Consume(line []byte) error {
	return nil
}

func (h *Part1Handler) Answer() int {
	return 0
}

func (h *Part1Handler) Print() {
	fmt.Printf("???")
}
