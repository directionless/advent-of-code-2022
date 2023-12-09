package day08

import (
	"errors"
	"fmt"
	"math/big"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	ExampleAnswer1 = 6
	ExampleAnswer2 = 6

	RealAnswer1 = 12737
	// 3,700,000,000 too low (max single thread brute force)
	// 4,619,677,975,437,858,519 is too high
	// 33,449,997,431 nope
	RealAnswer2 = 9064949303801
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

// findPeriod find the number of steps in a loop
func (h *dayHandler) findPeriod(loc locationType) (int, error) {
	if len(h.lrInstructions) == 0 {
		return -1, errors.New("lrInstructions is unset")
	}

	fmt.Printf("Finding period for %s\n", loc.Name)

	startingLoc := loc
	numberOfSteps := 0
	lastStep := 0
	lastZ := startingLoc.Name
	lastDelta := 0

	for {
		s := numberOfSteps % len(h.lrInstructions)
		numberOfSteps += 1

		if numberOfSteps > 1_000_000 {
			return -1, errors.New("unable to find loop: too many steps")
		}

		switch h.lrInstructions[s] {
		case byte('L'):
			loc = h.desertMap[loc.L]
		case byte('R'):
			loc = h.desertMap[loc.R]
		default:
			fmt.Printf("Unknown step direction %s\n", string(h.lrInstructions[s]))
			panic("unknown step direction")
		}

		if loc.GhostZ() {
			stepDelta := numberOfSteps - lastStep
			lastStep = numberOfSteps

			fmt.Printf("Found a z %s after %d steps (delta %d)\n", loc.Name, numberOfSteps, stepDelta)

			if loc.Equal(lastZ) && lastDelta == stepDelta {
				// confirmed a loop
				return stepDelta, nil
			}

			lastZ = loc.Name
			lastDelta = stepDelta
		}
	}
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

func (h *dayHandler) AnswerPart2() int {
	locations := map[locationType]int{}

	for name, loc := range h.desertMap {
		if name[2] == byte('A') {
			fmt.Printf("Starting location %s. Looking for period\n", loc.Name)
			period, err := h.findPeriod(loc)
			if err != nil {
				panic(err)
			}
			locations[loc] = period
		}
	}

	lcm := 1

	p := message.NewPrinter(language.English)
	for loc, period := range locations {
		cycleCount := float64(period) / float64(len(h.lrInstructions))

		p.Printf("location %s has a period of %d maybe factor %f (prime?) %v\n",
			loc.Name,
			period,
			cycleCount,
			big.NewInt(int64(cycleCount)).ProbablyPrime(0))

		lcm = LCM(lcm, period)
	}

	fmt.Printf("They should overlap after %d steps\n", lcm)
	return lcm
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
