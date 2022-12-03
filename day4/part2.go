package day4

import "fmt"

type Part2Handler struct {
}

func NewPart2() *Part2Handler {
	return &Part2Handler{}
}

func (h *Part2Handler) Consume(line []byte) error {
	return nil
}

func (h *Part2Handler) Answer() int {
	return 0
}

func (h *Part2Handler) Print() {
	fmt.Printf("???")
}
