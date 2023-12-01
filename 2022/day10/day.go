package day10

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/directionless/advent-of-code-2022/day10/cpu"
	"github.com/directionless/advent-of-code-2022/day10/crt"
)

type part1Tap struct {
	sum int
}

func (t *part1Tap) Examine(cycle int, x int) {
	switch cycle {
	case 20, 60, 100, 140, 180, 220:
		strength := x * cycle
		t.sum += strength

		fmt.Printf("cycle%d: x:%d, strength:%d\n", cycle, x, strength)
	}
}

func (t *part1Tap) Sum() int {
	return t.sum
}

type dayHandler struct {
	cpu         *cpu.CPU
	finalTapRun bool
	part1       *part1Tap
	crt         *crt.CrtStruct
}

func New() *dayHandler {
	h := &dayHandler{
		cpu:   cpu.New(),
		part1: &part1Tap{},
		crt:   crt.New(40, 6),
	}

	h.cpu.AddTap(h.part1)
	h.cpu.AddTap(h.crt)
	return h

}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	switch {
	case bytes.HasPrefix(line, []byte("noop")):
		h.cpu.ExecNoop()
	case bytes.HasPrefix(line, []byte("addx")):
		rawX := line[5:]
		x, err := strconv.Atoi(string(rawX))
		if err != nil {
			return fmt.Errorf("unable to parse x: %w", err)
		}
		h.cpu.ExecAddx(x)
	default:
		return fmt.Errorf("unknown instruction: %v", line)
	}
	return nil
}

func (h *dayHandler) AnswerPart1() int {
	// Note that the taps trigger after a tick, and depending on the input there might be
	// a fencepost. Practially speaking, the input seems to have a lot of trailing noops
	// to avoid this problem.
	return h.part1.Sum()
}

func (h *dayHandler) GetCRT(f io.Writer) {
	h.crt.OutputF(f)
}

func (h *dayHandler) AnswerPart2() int {
	// Note that the taps trigger after a tick, and depending on the input there might be
	// a fencepost. Practially speaking, the input seems to have a lot of trailing noops
	// to avoid this problem.

	h.crt.Output()
	return 0
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
