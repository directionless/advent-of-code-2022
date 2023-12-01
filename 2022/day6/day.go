package day6

import (
	"bufio"
	"fmt"
	"io"
)

type dayHandler struct {
	charactersProcessedToFindMarker  int
	charactersProcessedToFindMessage int
}

func New() *dayHandler {
	return &dayHandler{}
}

func (h *dayHandler) Handle(in io.Reader) error {
	//fmt.Printf("line: %s\n", string(line))

	startOfMarker := NewPacket(4)
	message := NewPacket(14)
	foundMarker := false

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanBytes)
	charNum := 0
	for scanner.Scan() {
		// these are human numbers, not indexes. So increment first
		charNum++

		b := scanner.Bytes()[0]

		//fmt.Printf("char %d, packet %s\n", charNum, message.String())

		if !foundMarker {
			startOfMarker.Push(b)
			if startOfMarker.Uniq() {
				h.charactersProcessedToFindMarker = charNum
				foundMarker = true
			}
		}

		message.Push(b)
		if message.Uniq() {
			h.charactersProcessedToFindMessage = charNum
			break
		}

	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan: %w", err)
	}

	return nil
}

func (h *dayHandler) AnswerPart1() int {
	return h.charactersProcessedToFindMarker
}

func (h *dayHandler) AnswerPart2() int {
	return h.charactersProcessedToFindMessage
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: Characters Processed to find marker: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: Characters processed to find message: %d\n", h.AnswerPart2())
}
