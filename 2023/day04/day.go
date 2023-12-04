package day04

import (
	"fmt"
	"regexp"
	"strings"
)

type dayHandler struct {
	part1_value int
}

func New() *dayHandler {
	h := &dayHandler{}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	cardValue, err := cardValue(string(line))
	if err != nil {
		return fmt.Errorf(`processing "%s": %w`, line, err)
	}

	h.part1_value += cardValue

	return nil
}

func (h *dayHandler) Solve() error {
	return nil
}

func (h *dayHandler) AnswerPart1() any {
	return h.part1_value
}

func (h *dayHandler) AnswerPart2() any {
	return nil
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}

var numRE = regexp.MustCompile(`([0-9]+)`)

func cardValue(line string) (int, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("expected 2 parts, got %d", len(parts))
	}

	card := parts[1]

	delimInd := strings.IndexRune(card, '|')
	myNumbers := numRE.FindAllString(card[:delimInd], -1)
	winNumbers := numRE.FindAllString(card[delimInd:], -1)

	// Need to find the intersection between myNumbers and the winNumbers. Golang doesn't have a
	// native intersection. One way would be to toss everything into a map[string]bool and then
	// iterate the second looking. Another way would be to call strings.Contains once per item
	// in a slice. I'm not sure which is faster. Either way though, myNumbers has fewer entries
	matches := 0
	matchHash := make(map[string]bool)
	for _, n := range winNumbers {
		matchHash[n] = true
	}
	for _, n := range myNumbers {
		if _, ok := matchHash[n]; ok {
			matches += 1
		}
	}

	fmt.Printf("Solving\n%s\n", line)
	fmt.Printf("win: %v\n my: %v\n", winNumbers, myNumbers)
	fmt.Println(parts[0], "num matches:", matches)
	// go doesn't have a quick power function (math is all floaty). Since there are at most
	// 10 matches we're going to be really sloppy and hardcode.
	if matches == 0 {
		return 0, nil
	}

	if matches > 10 {
		return 0, fmt.Errorf("too many matches. Had %d", matches)
	}

	return 1 << (matches - 1), nil

}
