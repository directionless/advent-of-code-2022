package day08

import (
	"fmt"
)

const (
	ExampleAnswer1 = -1
	ExampleAnswer2 = -1

	RealAnswer1 = 12737
	RealAnswer2 = -1
)

type dayHandler struct {
	lrInstructions []byte
	desertMap      map[[3]byte]locationType

	part1Answer any
	part2Answer any
}

func New() *dayHandler {
	h := &dayHandler{
		desertMap: make(map[[3]byte]locationType, 0),
	}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	// hack to grab the first line
	if len(h.lrInstructions) == 0 {
		h.lrInstructions = line
		return nil
	}

	loc, err := locationFromLine(line)
	if err != nil {
		return err
	}

	h.desertMap[loc.Name] = loc
	return nil
}

// Solve is called when the input is done being Consumed. Some puzzle can be solved entirely
// in Consume, line by line. Others need an additional step
func (h *dayHandler) Solve() error {
	// This has the handy effect of catchig the EOF. Advance lines, and call solve.
	return nil
}

func (h *dayHandler) AnswerPart1() any {

	for _, loc := range h.desertMap {
		fmt.Println(loc.DebugString())
	}

	loc, ok := h.desertMap[[3]byte{'A', 'A', 'A'}]
	if !ok {
		panic("No starting location")
	}

	numberOfSteps := 0
	for {
		fmt.Printf("In node %s\n", loc.Name)
		s := numberOfSteps % len(h.lrInstructions)
		switch h.lrInstructions[s] {
		case byte('L'):
			loc = h.desertMap[loc.L]
		case byte('R'):
			loc = h.desertMap[loc.R]
		default:
			fmt.Printf("Unknown step direction %s\n", string(h.lrInstructions[s]))
			panic("unknown step direction")
		}

		numberOfSteps += 1
		if loc.ZZZ() {
			break
		}
	}

	return numberOfSteps

}

func (h *dayHandler) AnswerPart2() any {
	return h.part2Answer

}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
