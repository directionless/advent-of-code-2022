package day16

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

const ()

type dayHandler struct {
	n *network
}

func New() *dayHandler {
	h := &dayHandler{
		n: newNetwork(),
	}

	return h

}

// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
var lineRe = regexp.MustCompile(`^Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)

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

	if len(m[0]) != 4 {
		spew.Dump(m[0])
		return fmt.Errorf("unexpected line: %s. got %d matches", line, len(m[0]))
	}

	valve := string(m[0][1])
	flowRate, err := strconv.Atoi(string(m[0][2]))
	if err != nil {
		return fmt.Errorf("unexpected line: %s. flow rate not an int: %s", line, err)
	}

	tunnels := bytes.Split(m[0][3], []byte(", "))

	h.n.AddValve(valve, flowRate, tunnels)

	return nil
}

func (h *dayHandler) AnswerPart1() int {
	pt, err := h.n.SolvePart1()
	if err != nil {
		panic(err)
	}

	spew.Dump(pt)
	return pt.MP
}

func (h *dayHandler) AnswerPart2() int {
	return 0
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
