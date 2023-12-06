package day05

import (
	"bytes"
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

	part1Answer int
	part2Answer int

	almanacMaps [][]thingThatNeedsMap
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
		seeds := extract.NumbersFromLine(line)
		h.seeds = make([]int, len(seeds))
		h.runningSeedValues = make(map[int]int, len(seeds))
		for i, s := range seeds {
			h.seeds[i] = s.Int
			h.runningSeedValues[s.Int] = s.Int
		}

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
	h.almanacMaps = append(h.almanacMaps, h.currentMap)
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
	lowest := -1

	for _, val := range h.seeds {
		seed := val
		for _, almanacMap := range h.almanacMaps {
			//fmt.Printf("testing almanac that has %s\n", almanacMap[0])
		NextAlmanac:
			for _, thing := range almanacMap {
				//fmt.Printf("seed %d now testing thing %d: %s\n", seed, i, thing)
				if newVal := thing.Contains(val); newVal >= 0 {
					fmt.Printf("seed %d is in thing: %s. Value %d -> %d\n", seed, thing, val, newVal)
					val = newVal
					break NextAlmanac
				}
			}
		}

		if val < lowest || lowest == -1 {
			fmt.Printf("seed %d has new lowest %d\n", seed, val)
			lowest = val
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
