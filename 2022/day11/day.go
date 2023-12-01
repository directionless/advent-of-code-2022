package day11

import (
	"bytes"
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

const ()

type dayHandler struct {
	curMonkey         *Monkey
	monkeyNet         *monkeyNetwork
	noRelief          bool
	commonWorryFactor int
}

func New() *dayHandler {
	h := &dayHandler{
		monkeyNet:         NewMonkeyNetwork(),
		commonWorryFactor: 1,
	}

	return h

}

// Consume consumes input. In the case of the monkeys, the input is multiline, but it's
// just as easy to parse line by line and keep state.
func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	switch {
	case bytes.HasPrefix(line, []byte("Monkey ")):
		m, err := strconv.Atoi(string(line[7 : len(line)-1]))
		if err != nil {
			return fmt.Errorf("unable to parse monkey number: %w", err)
		}

		h.curMonkey = NewMonkey(m, h.noRelief)
		h.monkeyNet.AddMonkey(h.curMonkey)

	case bytes.HasPrefix(line, []byte("  Starting items: ")):
		items := bytes.Split(line[18:], []byte(", "))
		for _, item := range items {
			w, err := strconv.Atoi(string(item))
			if err != nil {
				return fmt.Errorf("unable to parse item: %w", err)
			}
			h.curMonkey.Push(NewItem(w))
		}

	case bytes.HasPrefix(line, []byte("  Operation: new = old ")):
		args := bytes.Split(line[23:], []byte(" "))

		if len(args) != 2 {
			return fmt.Errorf("unable to parse operation from line: %s", line)
		}

		if slices.Equal(args[0], []byte("*")) && slices.Equal(args[1], []byte("old")) {
			h.curMonkey.Inspect = genOperateSquare()
			return nil
		}

		num, err := strconv.Atoi(string(args[1]))
		if err != nil {
			return fmt.Errorf("unable to parse number: %w", err)
		}

		switch {
		case slices.Equal(args[0], []byte("+")):
			h.curMonkey.Inspect = genOperateAdd(num)
		case slices.Equal(args[0], []byte("*")):
			h.curMonkey.Inspect = genOperateMultiply(num)
		default:
			return fmt.Errorf("unknown operand: %s", string(args[0]))
		}

	case bytes.HasPrefix(line, []byte("  Test: ")):
		args := bytes.Split(line[8:], []byte(" by "))
		if len(args) != 2 {
			return fmt.Errorf("unable to parse test from line: %s", line)
		}

		if !slices.Equal(args[0], []byte("divisible")) {
			return fmt.Errorf("don't know how to make a test for %s", args[0])
		}

		num, err := strconv.Atoi(string(args[1]))
		if err != nil {
			return fmt.Errorf("unable to parse number: %w", err)
		}

		h.monkeyNet.commonWorryFactor *= num

		h.curMonkey.TestFn = genTestDivisible(num)
	case bytes.HasPrefix(line, []byte("    If true: throw to monkey ")):
		n, err := strconv.Atoi(string(line[29:]))
		if err != nil {
			return fmt.Errorf("unable to parse monkey number: %w", err)
		}
		h.curMonkey.TrueDest = n

	case bytes.HasPrefix(line, []byte("    If false: throw to monkey ")):
		n, err := strconv.Atoi(string(line[30:]))
		if err != nil {
			return fmt.Errorf("unable to parse monkey number: %w", err)
		}
		h.curMonkey.FalseDest = n

	default:
		return fmt.Errorf("unable to parse line: %s", line)
	}
	return nil
}

func (h *dayHandler) Run(rounds int) error {
	for i := 1; i <= rounds; i++ {
		//fmt.Printf("Starting Round %d\n", i)
		if err := h.monkeyNet.RunRound(); err != nil {
			return fmt.Errorf("unable to run round %d: %w", i+1, err)
		}

		if i == 1 || i == 20 || i == 1000 {
			fmt.Printf("After Round %d\n", i)
			h.monkeyNet.PrintInfo()
		}
	}
	return nil
}

func (h *dayHandler) AnswerPart1() int {
	h.monkeyNet.PrintInfo()
	return h.monkeyNet.MonkeyBusiness()
}

func (h *dayHandler) AnswerPart2() int {
	h.monkeyNet.PrintInfo()
	return h.monkeyNet.MonkeyBusiness()
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
