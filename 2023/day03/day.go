package day03

import (
	"fmt"
)

type dayHandler struct {
	prevLine    *lineInfo
	currentLine *lineInfo
	nextLine    *lineInfo

	part1_total int
	part2_total int
}

func New() *dayHandler {
	h := &dayHandler{
		prevLine:    emptyLine,
		currentLine: emptyLine,
		nextLine:    emptyLine,
	}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	li, err := parseLine(string(line))
	if err != nil {
		return fmt.Errorf(`parsing line "%s": %w`, line, err)
	}

	// As this starts a new line, we need to advance our various lines
	// There's probably an optimization to be had here, around _not_ reallocating buffers.
	// But I'm not sure I'm clever enough to find it, especially since the elves sometime
	// change line widths mid file. So, we bail to strings, it's small enough the extra
	// allocations probably don't matter. (And if the compiler is clever enough, there won't
	// be extra allocations, this could all be pointer esque)
	h.prevLine = h.currentLine
	h.currentLine = h.nextLine
	h.nextLine = li

	matchingNums, gearRatios, err := solveLine(h.prevLine, h.currentLine, h.nextLine)
	if err != nil {
		return fmt.Errorf(`solving line "%s": %w`, line, err)
	}

	for _, n := range matchingNums {
		h.part1_total += n
	}

	h.part2_total += gearRatios

	return nil
}

func (h *dayHandler) Solve() error {
	// This has the handy effect of catchig the EOF. Advance lines, and call solve.
	h.prevLine = h.currentLine
	h.currentLine = h.nextLine
	h.nextLine = emptyLine

	matchingNums, gearRatios, err := solveLine(h.prevLine, h.currentLine, h.nextLine)
	if err != nil {
		return fmt.Errorf(`solving EOF: %w`, err)
	}

	for _, n := range matchingNums {
		h.part1_total += n
	}

	h.part2_total += gearRatios

	return nil
}

// solveLine acts on the current 3 lines
func solveLine(prevLine, curLine, nextLine *lineInfo) ([]int, int, error) {
	// The first time we're called, we're not actually ready to solve. So early return.
	if curLine.Empty() {
		return nil, 0, nil
	}
	fmt.Printf("\n\nSolving\n  %s\n> %s\n  %s\n\n", prevLine, curLine, nextLine)

	symbolIndexes := []int{}
	for _, li := range []*lineInfo{prevLine, curLine, nextLine} {
		symbolIndexes = append(symbolIndexes, li.SymbolPositions()...)
	}

	matches := curLine.NumbersTouching(symbolIndexes)
	//fmt.Printf("matching numbers: %v\n", matches)

	// Part2
	//
	// Lets fine the gears. This is slightly inefficient, because we end up calling `NumbersTouching`
	// twice. But, :shrug: short enough
	gearRatios := 0
	fmt.Printf("Gear Index: %v\n", curLine.GearIndexes())

	for _, gearIdx := range curLine.GearIndexes() {
		nearNumbers := []int{}
		for _, li := range []*lineInfo{prevLine, curLine, nextLine} {
			nearNumbers = append(nearNumbers, li.NumbersTouching([]int{gearIdx})...)
		}

		if len(nearNumbers) == 2 {
			gearRatio := nearNumbers[0] * nearNumbers[1]
			fmt.Printf("Found gear: idx %d: %d and %d -- ratio %d\n", gearIdx, nearNumbers[0], nearNumbers[1], gearRatio)
			gearRatios += gearRatio
		}

	}

	return matches, gearRatios, nil
}

func (h *dayHandler) AnswerPart1() any {
	return h.part1_total
}

func (h *dayHandler) AnswerPart2() any {
	return h.part2_total
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
