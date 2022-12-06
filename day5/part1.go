package day5

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

// Simple enough move extractor
var moveRegexp = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

type Part1Handler struct {
	rawStacks      map[int][]byte
	stacks         map[int][]byte
	stackNames     map[int]string
	stackNameToIdx map[byte]int
	doneWithHeader bool
}

func New() *Part1Handler {
	return &Part1Handler{
		rawStacks: make(map[int][]byte, 9),
	}
}

func (h *Part1Handler) Consume(line []byte) error {
	// Since we don't have access to the iterator (abstraction oops!) we can't just call
	// next several times. Instead, we need to splay over a couple state vars.
	if len(line) == 0 {
		h.doneWithHeader = true
		if err := h.stackReverser(); err != nil {
			return fmt.Errorf("unable to reverse stacks: %v", err)
		}
	}

	if !h.doneWithHeader {
		if err := h.columnParser(line); err != nil {
			return fmt.Errorf("unable to parse %s: %w", line, err)
		}
	}

	// Should be some move commands
	m := moveRegexp.FindAllSubmatch(line, -1)
	if m != nil {
		count, err := strconv.Atoi(string(m[0][1]))
		if err != nil {
			return fmt.Errorf("unable to discern count from %s on line %s: %w", m[0][1], line, err)
		}

		from := m[0][2]
		to := m[0][3]

		if err := h.operateCrane9001(count, from, to); err != nil {
			return fmt.Errorf("failed to operate crane for line: %s: %w", line, err)
		}

	}

	return nil
}

func (h *Part1Handler) AnswerPart1() int {
	names := []byte{}

	for n := range h.stackNameToIdx {
		names = append(names, n)
	}

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	answer := []byte{}
	for _, n := range names {
		stackIdx := h.stackNameToIdx[n]
		topItem := h.stacks[stackIdx][len(h.stacks[stackIdx])-1]

		answer = append(answer, topItem)

		fmt.Printf("Stack %s top item %s\n", string(n), string(topItem))
	}

	fmt.Printf("Part1 Answer is %s\n", answer)
	return 0
}

func (h *Part1Handler) AnswerPart2() int {
	return 0
}

func (h *Part1Handler) Print() {
	fmt.Printf("???")
}

func (h *Part1Handler) PrintStacks() error {
	if len(h.stacks) != len(h.stackNames) {
		return errors.New("stacks and stack names are not the same length")
	}

	for i, stack := range h.stacks {
		fmt.Fprintf(os.Stderr, "%s: %s\n", h.stackNames[i], stack)
	}
	return nil
}

func (h Part1Handler) operateCrane9000(count int, fromName, toName []byte) error {
	fromIdx, ok := h.stackNameToIdx[fromName[0]]
	if !ok {
		return fmt.Errorf("from stack %s not found", fromName)
	}

	toIdx, ok := h.stackNameToIdx[toName[0]]
	if !ok {
		return fmt.Errorf("to stack %s not found", toName)
	}

	fromLen := len(h.stacks[fromIdx])
	if fromLen < count {
		return fmt.Errorf("stack %s doesn't have %d items left. Only %d", fromName, count, fromLen)
	}

	for i := 0; i < count; i++ {
		fromLen = len(h.stacks[fromIdx])

		item := h.stacks[fromIdx][fromLen-1]
		h.stacks[fromIdx] = h.stacks[fromIdx][:fromLen-1]

		h.stacks[toIdx] = append(h.stacks[toIdx], item)
	}

	return nil
}

func (h Part1Handler) operateCrane9001(count int, fromName, toName []byte) error {
	fromIdx, ok := h.stackNameToIdx[fromName[0]]
	if !ok {
		return fmt.Errorf("from stack %s not found", fromName)
	}

	toIdx, ok := h.stackNameToIdx[toName[0]]
	if !ok {
		return fmt.Errorf("to stack %s not found", toName)
	}

	fromLen := len(h.stacks[fromIdx])
	if fromLen < count {
		return fmt.Errorf("stack %s doesn't have %d items left. Only %d", fromName, count, fromLen)
	}

	//fmt.Printf("Attempting to take %d items off %v\n", count, h.stacks[fromIdx])
	items := h.stacks[fromIdx][fromLen-count : fromLen]
	//fmt.Printf("Got %v\n", items)
	h.stacks[toIdx] = append(h.stacks[toIdx], items...)

	h.stacks[fromIdx] = h.stacks[fromIdx][:fromLen-count]

	return nil
}

func (h *Part1Handler) stackReverser() error {
	h.stacks = make(map[int][]byte, len(h.rawStacks))
	h.stackNames = make(map[int]string, len(h.rawStacks))
	h.stackNameToIdx = make(map[byte]int, len(h.rawStacks))

	for stackIdx, stack := range h.rawStacks {
		// The last item on the stack is the name.
		stackName := stack[len(stack)-1]
		h.stackNames[stackIdx] = string(stackName)
		h.stackNameToIdx[stackName] = stackIdx

		// The rest of the items are the actual items.

		// The rest of the stack is the actual stack.
		h.stacks[stackIdx] = make([]byte, len(stack)-1)
		newPos := 0
		for i := len(stack) - 2; i >= 0; i-- {
			h.stacks[stackIdx][newPos] = stack[i]
			newPos++
		}
	}

	return nil
}

func (h *Part1Handler) columnParser(line []byte) error {
	// we don't know how many columns, but we don't need to. Columns take up 3 spaces, with a space seperator. That
	// means we can consume 4 characters at a time and see when we run out.

	if len(line) < 4 {
		return fmt.Errorf("line too short: %s", line)
	}

	if line[0] != []byte("[")[0] && line[0] != []byte(" ")[0] {
		return fmt.Errorf("unexpected line start: %s", string(line[0]))
	}

	for i, b := range line {
		pos := i % 4
		stackIdx := i / 4

		switch pos {
		case 0:
			if b != []byte(`[`)[0] && b != []byte(" ")[0] {
				return fmt.Errorf("character %s is a surprise. expected open", string(b))
			}
		case 1:
			if b == []byte(" ")[0] {
				continue
			}

			if h.rawStacks[stackIdx] == nil {
				h.rawStacks[stackIdx] = []byte{b}
			} else {
				h.rawStacks[stackIdx] = append(h.rawStacks[stackIdx], b)
			}
			// fmt.Printf("%d(%s)", stackIdx, string(b))
		case 2:
			if b != []byte("]")[0] && b != []byte(" ")[0] {
				return fmt.Errorf("character %s is a surprise. expected close", string(b))
			}
		case 3:
			if b != []byte(" ")[0] {
				return fmt.Errorf("character %s is a surprise. expected seperator", string(b))
			}
		}
	}

	return nil
}

func numColumns(line []byte) int {
	// Each column is represented by 3 characters, plus a space seperator.
	// Therefore, the line length is 4 * columns -1
	return (len(line) + 1) / 4
}
