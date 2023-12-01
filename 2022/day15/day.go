package day15

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

const ()

type dayHandler struct {
	grid     *grid
	network  *network
	part1Row int
	part2Max int
}

func New(part1Row int, part2Max int) *dayHandler {
	h := &dayHandler{
		grid:     NewGrid(),
		network:  NewNetwork(),
		part1Row: part1Row,
		part2Max: part2Max,
	}

	return h

}

var lineRe = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	if line[0] == '#' {
		return nil
	}

	m := lineRe.FindAllSubmatch(line, -1)
	if m == nil || len(m) != 1 {
		return fmt.Errorf("unexpected line: %s. did not match", line)
	}

	if len(m[0]) != 5 {
		spew.Dump(m[0])
		return fmt.Errorf("unexpected line: %s. got %d matches", line, len(m[0]))
	}

	asInt := make([]int, 5)
	for i := 1; i < 5; i++ {
		var err error
		asInt[i], err = strconv.Atoi(string(m[0][i]))
		if err != nil {
			return fmt.Errorf("could not parse int: %w", err)
		}
	}

	sensor := location{X: asInt[1], Y: asInt[2]}
	beacon := location{X: asInt[3], Y: asInt[4]}

	//fmt.Printf("Beacon: %s, snsor: %s\n", beacon, sensor)

	// Grid is too inefficient to use in real. But the pretty print is nice
	//  for debugging examples
	//h.grid.AddSensor(sensor, beacon)

	h.network.AddDetection(sensor, beacon)

	return nil
}

func (h *dayHandler) AnswerPart1() int {
	count := h.network.GetRowY(h.part1Row)
	//fmt.Println(pretty)
	return count
}

func (h *dayHandler) AnswerPart2() int {
	//fmt.Println(h.grid)

	loc, err := h.network.FindHoles(h.part2Max)
	if err != nil {
		panic(err)
	}

	return loc.X*4000000 + loc.Y
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
