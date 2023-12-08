package day06

import (
	"fmt"
	"strings"

	"github.com/directionless/advent-of-code-2022/2023/util/extract"
)

const (
	ExampleAnswer1 = 288
	ExampleAnswer2 = 71503

	RealAnswer1 = 316800
	RealAnswer2 = 45647654
)

type dayHandler struct {
	times     []int
	distances []int

	time2     int
	distance2 int

	part1Answer int
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

	switch {
	case strings.HasPrefix(string(line), "Time"):
		times := extract.NumbersFromLine(line)
		h.times = make([]int, len(times))
		for i, s := range times {
			h.times[i] = s.Int
		}

		part2 := extract.NumbersFromLine([]byte(strings.Replace(string(line), " ", "", -1)))
		h.time2 = part2[0].Int

	case strings.HasPrefix(string(line), "Distance"):
		distances := extract.NumbersFromLine(line)
		h.distances = make([]int, len(distances))
		for i, s := range distances {
			h.distances[i] = s.Int
		}

		part2 := extract.NumbersFromLine([]byte(strings.Replace(string(line), " ", "", -1)))
		h.distance2 = part2[0].Int

	default:
		return fmt.Errorf("unknown line: %s", line)
	}
	return nil
}

// Solve is called when the input is done being Consumed. Some puzzle can be solved entirely
// in Consume, line by line. Others need an additional step
func (h *dayHandler) Solve() error {
	// This has the handy effect of catchig the EOF. Advance lines, and call solve.
	if len(h.times) != len(h.distances) {
		return fmt.Errorf("length mismatch")
	}

	return nil
}

func findRaceResults(maxTime int) map[int]int {
	results := make(map[int]int, maxTime-1)
	for chargeDuration := 0; chargeDuration < maxTime; chargeDuration += 1 {
		results[chargeDuration] = (maxTime - chargeDuration) * chargeDuration
	}
	return results
}

/*
func findWinningRaceResults(maxTime, winDistance int) map[int]int {
	waysToWin := 0

	for chargeDuration := 0; chargeDuration < maxTime; chargeDuration += 1 {
		results[chargeDuration] = (maxTime - chargeDuration) * chargeDuration
	}
	return results
}
*/

func (h *dayHandler) AnswerPart1() any {
	h.part1Answer = 1
	for i, maxTime := range h.times {
		winDistance := h.distances[i]
		waysToWin := 0
		possibleRaces := findRaceResults(maxTime)
		for time, distance := range possibleRaces {
			_ = time
			// fmt.Printf("total time %d: charge for %d distance traveled %d\n", maxTime, time, distance)
			if distance > winDistance {
				waysToWin += 1
			}
		}

		fmt.Printf("Race %d has %d ways to win\n", maxTime, waysToWin)

		h.part1Answer *= waysToWin
	}

	return h.part1Answer

}

func (h *dayHandler) AnswerPart2() any {
	waysToWin := 0
	possibleRaces := findRaceResults(h.time2)
	for time, distance := range possibleRaces {
		_ = time
		// fmt.Printf("total time %d: charge for %d distance traveled %d\n", maxTime, time, distance)
		if distance > h.distance2 {
			waysToWin += 1
		}
	}

	h.part2Answer = waysToWin

	return h.part2Answer

}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
