package day21

import (
	"fmt"
	"regexp"
	"strconv"
)

const ()

// Examples:
//   sjmn: drzm * dbpl
//   sllz: 4

var monkeyRe = regexp.MustCompile(`(....): (?:(\d+)|(....) (.) (....))$`)

type dayHandler struct {
	monkeys map[string]monkeyRun
}

func New() *dayHandler {
	h := &dayHandler{
		monkeys: make(map[string]monkeyRun),
	}
	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	m := monkeyRe.FindAllSubmatch(line, -1)
	if m == nil {
		return fmt.Errorf("unexpected line: %s", line)
	}

	monkey := string(m[0][1])
	con := string(m[0][2])
	a := string(m[0][3])
	op := string(m[0][4])
	b := string(m[0][5])

	if len(con) > 0 {
		num, err := strconv.Atoi(con)
		if err != nil {
			return fmt.Errorf("converting line %s: %w", line, err)
		}
		h.monkeys[monkey] = &monkeyConstant{num}
		return nil
	}

	switch op {
	case "+":
		h.monkeys[monkey] = &monkeyAdd{a, b}
	case "-":
		h.monkeys[monkey] = &monkeySub{a, b}
	case "*":
		h.monkeys[monkey] = &monkeyMul{a, b}
	case "/":
		h.monkeys[monkey] = &monkeyDiv{a, b}
	}

	return nil
}

func (h *dayHandler) AnswerPart1() int {
	answer, err := h.monkeys["root"].Run(h.monkeys)
	if err != nil {
		panic(err)
	}

	return answer
}

func (h *dayHandler) AnswerPart2() int {
	return 0
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
