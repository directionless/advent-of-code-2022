package day20

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const ()

type inputNum struct {
	num int
}

func (i *inputNum) String() string {
	return fmt.Sprintf("%d", i.num)
}

func (i *inputNum) N() int {
	return i.num
}

type dayHandler struct {
	ring    *list.List
	initial []*list.Element
	zero    *list.Element
}

func New() *dayHandler {
	h := &dayHandler{
		ring:    list.New(),
		initial: make([]*list.Element, 0),
	}
	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	num, err := strconv.Atoi(string(line))
	if err != nil {
		return fmt.Errorf("unable to parse line %s: %w", line, err)
	}

	in := &inputNum{
		num: num,
	}

	e := h.ring.PushBack(in)
	h.initial = append(h.initial, e)

	return nil
}

func (h *dayHandler) Solve() error {
	for i, e := range h.initial {
		h.PrintRing()
		fmt.Printf("\nNow moving %s (step %d)\n", e.Value, i)
		h.MoveElement(e)
	}
	h.PrintRing()

	h.zero = h.FindZero()
	if h.zero == nil {
		return errors.New("unable to find 0")
	}

	return nil
}

func (h *dayHandler) FindZero() *list.Element {
	for e := h.ring.Front(); e != nil; e = e.Next() {
		if e.Value.(*inputNum).num == 0 {
			return e
		}
	}
	return nil
}

func (h *dayHandler) ElementNextPos(e *list.Element, num int) *list.Element {
	switch {
	case num >= 1:
		newPos := e.Next()
		if newPos == nil {
			newPos = h.ring.Front() // wrap around

		}
		return newPos
	case num <= -1:
		newPos := e.Prev()
		if newPos == nil {
			newPos = h.ring.Back() // wrap around
		}
		return newPos
	default:
		return nil
	}
}

func (h *dayHandler) MoveElement(e *list.Element) {
	num := e.Value.(*inputNum).N()

	// make n smaller, It gets weird if N wraps past current position.
	// the -1 is becasue we don't count this element while we move it
	num = (h.ring.Len() - 1) % num

	switch {
	case num == 0:
		return
	case num >= 1:
		var newPos *list.Element
		for i := 0; i < num; i++ {
			newPos = h.ElementNextPos(e, 1)
		}

		fmt.Printf("Moving %s after %s\n", e.Value, newPos.Value)
		h.ring.MoveAfter(e, newPos)

	case num < 1:
		var newPos *list.Element
		for i := 0; i < num; i++ {
			newPos = h.ElementNextPos(e, -1)
		}

		fmt.Printf("Moving %s before %s\n", e.Value, newPos.Value)
		h.ring.MoveBefore(e, newPos)
	}
}

func (h *dayHandler) PrintRing() {
	e := h.ring.Front()
	for e != nil {
		fmt.Printf("%s, ", e.Value)
		e = e.Next()
	}
	fmt.Println()
}

func (h *dayHandler) StringRing() string {
	sb := strings.Builder{}
	e := h.ring.Front()
	for e != nil {
		fmt.Fprintf(&sb, "%s, ", e.Value)
		e = e.Next()
	}
	return sb.String()
}

func (h *dayHandler) AnswerPart1() int {
	pos1000 := h.GetPos(1000)
	pos2000 := h.GetPos(2000)
	pos3000 := h.GetPos(3000)
	return pos1000.N() + pos2000.N() + pos3000.N()
}

func (h *dayHandler) GetPos(n int) *inputNum {
	e := h.zero

	for i := 0; i < n; i++ {
		e = e.Next()
		if e == nil {
			e = h.ring.Front()
		}
	}

	fmt.Printf("pos %d is %s\n", n, e.Value)

	return e.Value.(*inputNum)
}

func (h *dayHandler) AnswerPart2() int {
	return 0
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
