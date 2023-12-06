package day05

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/directionless/advent-of-code-2022/2023/util/extract"
)

const (
	ExampleAnswer1 = 35
	ExampleAnswer2 = 289863851

	RealAnswer1 = -1
	RealAnswer2 = -1
)

type dayHandler struct {
	seeds             []int
	currentMapName    string
	currentMap        []thingThatNeedsMap
	runningSeedValues map[int]int

	part1Answer any
	part2Answer any
}

func New() *dayHandler {
	h := &dayHandler{}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	// The maps show up in the order they're needed. At least for part 1. Yay
	if bytes.HasPrefix(line, []byte("seeds:")) {
		if len(h.seeds) != 0 {
			return errors.New("got a second seeds declaration")
		}

		seeds := extract.NumbersFromLine(line)
		h.seeds = make([]int, len(seeds))
		h.runningSeedValues = make(map[int]int, len(seeds))
		for i, s := range seeds {
			h.seeds[i] = s.Int
			h.runningSeedValues[s.Int] = s.Int
		}

		// This only handles single line of seeds, We could probably just drop this. But :shrug:
		return nil
	}

	if bytes.Contains(line, []byte(" map:")) {
		if err := h.processCompletedMap(); err != nil {
			return fmt.Errorf("processing map: %w", err)
		}

		parts := strings.SplitN(string(line), " ", 2)
		h.currentMapName = parts[0]
		h.currentMap = nil
		return nil
	}

	thing, err := thingFromLine(h.currentMapName, line)
	if err != nil {
		return fmt.Errorf("creating %s thing: %w", h.currentMapName, err)
	}

	//fmt.Printf("Found a thing!: %v\n", thing)
	h.currentMap = append(h.currentMap, thing)

	return nil
}

func (h *dayHandler) processCompletedMap() error {
	// if the map is empty, assume this is the first iterate and bail
	if len(h.currentMap) == 0 {
		return nil
	}

	fmt.Printf("currently processing map %s\n", h.currentMapName)
	sort.Sort(bySrcStart(h.currentMap))
	//spew.Dump(h.currentMap)
	for seed, val := range h.runningSeedValues {
		for _, thing := range h.currentMap {
			if newVal := thing.Contains(val); newVal >= 0 {
				//fmt.Printf("seed %d is in thing: %s. Value %d -> %d\n", seed, thing, val, newVal)
				h.runningSeedValues[seed] = newVal
			}
		}
	}

	return nil
}

// Solve is called when the input is done being Consumed. Some puzzle can be solved entirely
// in Consume, line by line. Others need an additional step
func (h *dayHandler) Solve() error {
	if err := h.processCompletedMap(); err != nil {
		return fmt.Errorf("processing map: %w", err)
	}

	return nil
}

func (h *dayHandler) AnswerPart1() any {
	// find the lowest seed value, which is location
	lowest := -1
	for seed, location := range h.runningSeedValues {
		fmt.Printf("seed %d is in location %d\n", seed, location)
		if location < lowest || lowest == -1 {
			lowest = location

		}
	}

	return lowest

}

func (h *dayHandler) AnswerPart2() any {
	return h.part2Answer

}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
