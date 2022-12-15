package day14

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/directionless/advent-of-code-2022/pkg/aocgrid"
)

const (
	sandStartX = 500
	sandStartY = 0
)

type dayHandler struct {
	grid        *aocgrid.Grid
	grains      int
	secretFloor int
	part2       bool
}

func New(part2 bool) *dayHandler {
	h := &dayHandler{
		grid:  aocgrid.New(),
		part2: part2,
	}

	h.grid.Set(sandStartX, sandStartY, SandSource)

	return h

}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	chain := bytes.Split(line, []byte(" -> "))
	for i, pt := range chain {
		x, y, err := parsePoint(pt)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", pt, err)
		}

		h.grid.Set(x, y, Rock)

		if i > 0 {
			// this causes extra parsing. But :shrug:
			x1, y1, err := parsePoint(chain[i-1])
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", chain[i-1], err)
			}

			x2, y2, err := parsePoint(chain[i])
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", chain[i], err)
			}

			h.grid.SetStraightLine(x1, y1, x2, y2, Rock)
		}
	}
	return nil
}

func (h *dayHandler) RunSand() error {
	h.grid.SetYMin(0)
	h.grid.FillRemaining(Air)

	// handling for part 2. Special y max!
	if h.part2 {
		h.secretFloor = h.grid.GetYMax() + 2
	}

	grains := 0
	for {
		grains += 1

		fmt.Printf("grain number: %d\n", grains)
		//fmt.Println(h.grid)

		if v := h.grid.Look(sandStartX, sandStartY); v != SandSource {
			break
		}

		if !h.AddSand(grains, sandStartX, sandStartY) {
			break
		}
	}

	fmt.Printf("grain number: %d\n", grains)
	fmt.Println(h.grid)

	// last grain didn't suceed, so omit it from the count
	h.grains = grains - 1

	return nil
}

// AddSand adds a grain of sand according to the rules. It returns false if sand fell
// into the void.
//
// A unit of sand always falls down one step if possible. If the tile immediately
// below is blocked (by rock or sand), the unit of sand attempts to instead move
// diagonally one step down and to the left. If that tile is blocked, the unit of
// sand attempts to instead move diagonally one step down and to the right. Sand
// keeps moving as long as it is able to do so, at each step trying to move down,
// then down-left, then down-right. If all three possible destinations are blocked,
// the unit of sand comes to rest and no longer moves, at which point the next unit
// of sand is created back at the source.
func (h *dayHandler) AddSand(grainNum, x, y int) bool {

	//fmt.Printf("grain number: %d\n", grainNum)
	//fmt.Printf("x: %d, y: %d\n", x, y)
	//fmt.Printf("down: %s, downleft: %s, downright: %s\n",
	//	down, downLeft, downRight)
	//fmt.Println(h.grid)

	// Part 2, expand the floor. This is somewhat inefficient -- we're adding a lot of
	// floor. But it's the quick way to fit this not-great idiom. So here we are.
	if h.part2 {
		h.grid.SetIfEmpty(x-1, h.secretFloor, Rock)
		h.grid.SetIfEmpty(x, h.secretFloor, Rock)
		h.grid.SetIfEmpty(x+1, h.secretFloor, Rock)
	}

	down := h.grid.LookDown(x, y)
	downLeft := h.grid.LookDownLeft(x, y)
	downRight := h.grid.LookDownRight(x, y)

	// Are we in the void? (part 1)
	if h.grid.Look(x, y) == nil && y > h.secretFloor {
		fmt.Println("IN VOID")
		return false
	}

	// Are we blocked by rock or sand?
	if (down == Rock || down == Sand) &&
		(downLeft == Rock || downLeft == Sand) &&
		(downRight == Rock || downRight == Sand) {
		h.grid.Set(x, y, Sand)
		return true
	}

	switch {
	case down == Air || down == nil:
		dx, dy := h.grid.Down(x, y)
		return h.AddSand(grainNum, dx, dy)

	case downLeft == Air || downLeft == nil:
		dx, dy := h.grid.DownLeft(x, y)
		return h.AddSand(grainNum, dx, dy)

	case downRight == Air || downRight == nil:
		dx, dy := h.grid.DownRight(x, y)
		return h.AddSand(grainNum, dx, dy)
	}

	// Shouldn't be possible to get here:
	fmt.Printf("SEPH: wtf %d,%d  down: %s, downLeft: %s, downRight: %s\n",
		x, y, down, downLeft, downRight)
	return false

}

func (h *dayHandler) AnswerPart1() int {
	return h.grains
}

func (h *dayHandler) AnswerPart2() int {
	return h.grains
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}

func parsePoint(raw []byte) (int, int, error) {
	nums := bytes.Split(raw, []byte(","))
	if len(nums) != 2 {
		return 0, 0, fmt.Errorf("invalid point: %s", raw)
	}

	x, err := strconv.Atoi(string(nums[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("unable to parse x: %w", err)
	}

	y, err := strconv.Atoi(string(nums[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("unable to parse y: %w", err)
	}

	return x, y, nil
}
