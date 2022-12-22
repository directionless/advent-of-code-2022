package day21

import (
	"errors"
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
	monkeys map[string]*monkeyTyp
}

func New() *dayHandler {
	h := &dayHandler{
		monkeys: make(map[string]*monkeyTyp),
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

	monkeyName := string(m[0][1])
	con := string(m[0][2])
	a := string(m[0][3])
	op := string(m[0][4])
	b := string(m[0][5])

	if len(con) > 0 {
		num, err := strconv.Atoi(con)
		if err != nil {
			return fmt.Errorf("converting line %s: %w", line, err)
		}
		h.findOrCreateMonkey(monkeyName).AddRun(&monkeyConstant{num})
	} else {

		switch op {
		case "+":
			h.findOrCreateMonkey(monkeyName).AddRun(&monkeyAdd{a, b})
			//h.findOrCreateMonkey(a).AddRun(&monkeySub{monkeyName, b})
			//h.findOrCreateMonkey(b).AddRun(&monkeySub{monkeyName, a})

		case "-":
			h.findOrCreateMonkey(monkeyName).AddRun(&monkeySub{a, b})
		case "*":
			h.findOrCreateMonkey(monkeyName).AddRun(&monkeyMul{a, b})
		case "/":
			h.findOrCreateMonkey(monkeyName).AddRun(&monkeyDiv{a, b})
		}

	}

	//h.monkeys[monkeyName] = monkey

	return nil
}

func (h *dayHandler) findOrCreateMonkey(name string) *monkeyTyp {
	_, ok := h.monkeys[name]
	if !ok {
		h.monkeys[name] = &monkeyTyp{}
	}

	return h.monkeys[name]
}

func (h *dayHandler) AnswerPart1() int {
	// dumb pointer / interface nonsense
	mm := make(monkeyMap, len(h.monkeys))
	for name, m := range h.monkeys {
		mm[name] = m
	}

	answer, err := h.monkeys["root"].Run(0, mm)
	if err != nil {
		panic(err)
	}

	return answer
}

func (h *dayHandler) part2Delta(mm monkeyMap, hum int, m1, m2 string) (int, error) {
	h.monkeys["humn"].OverrideHuman(hum)

	m1a, err := h.monkeys[m1].Run(0, mm)
	if err != nil {
		return 0, err
	}

	m2a, err := h.monkeys[m2].Run(0, mm)
	if err != nil {
		return 0, err
	}

	d2 := m2a - m1a
	return d2, nil
}

func (h *dayHandler) part2Check(mm monkeyMap, hum, d1 int, m1, m2 string) (int, error) {
	h.monkeys["humn"].OverrideHuman(hum)
	d2, err := h.part2Delta(mm, hum, m1, m2)
	if err != nil {
		return 0, err
	}

	fmt.Printf("human: %d, delta: %d, delta diff %d\n", hum, d2, d1-d2)

	switch {
	case d2 == 0:
		fmt.Println("FOUND")
		return hum, nil
	case d1 >= d2:
		return h.part2Check(mm, hum*2, d2, m1, m2)
	case d1 < d2:
		return h.part2Check(mm, hum/2, d2, m1, m2)
	}

	return 0, errors.New("wtf")
}

func (h *dayHandler) SearchPart2() error {
	mm := make(monkeyMap, len(h.monkeys))
	for name, m := range h.monkeys {
		mm[name] = m
	}

	m1, m2, err := h.monkeys["root"].ExtractRoot()
	if err != nil {
		return err
	}

	//startingHuman, err := h.monkeys["humn"].GetHuman()
	//if err != nil {
	//	return err
	//}

	// Horrible manual search. I changed constants until I found it
	for i := 3882224466100; i < 3882224466200; i = i + 1 {
		guess := i

		d1, err := h.part2Delta(mm, guess, m1, m2)
		if err != nil {
			return err
		}

		fmt.Printf("human: %d, delta: %d\n", guess, d1)
	}

	return nil
}

func (h *dayHandler) AnswerPart2() int {
	return 0
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
