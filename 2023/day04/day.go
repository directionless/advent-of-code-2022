package day04

import (
	"fmt"
	"regexp"
	"strings"
)

type dayHandler struct {
	part1_value       int
	idx               int
	part2_card_counts map[int]int
}

func New() *dayHandler {
	h := &dayHandler{
		part2_card_counts: make(map[int]int, 0),
	}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	// Increment _after_ we loop
	defer func() { h.idx += 1 }()

	// We always have at least _1_ copy of a scratch card.
	h.part2_card_counts[h.idx] += 1

	cardMatches, err := cardValue(string(line))
	if err != nil {
		return fmt.Errorf(`processing "%s": %w`, line, err)
	}

	h.part1_value += valueForMatches(cardMatches)

	// part2
	//
	// These elves, this is nonsense. So when a card has matches, it increases
	// the number of subsequent cards. This seems kinda bonkers. But hey...
	fmt.Printf("Card %d, has %d copies\n", h.idx+1, h.part2_card_counts[h.idx])
	for i := 1; i <= cardMatches; i++ {
		// Ever time a card has N matches, it increases the next N card counts by 1.
		h.part2_card_counts[h.idx+i] += h.part2_card_counts[h.idx]
	}

	return nil
}

func (h *dayHandler) Solve() error {
	return nil
}

func (h *dayHandler) AnswerPart1() any {
	return h.part1_value
}

func (h *dayHandler) AnswerPart2() any {
	sum := 0
	for _, n := range h.part2_card_counts {
		sum += n
	}
	return sum
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}

func valueForMatches(matches int) int {
	// go doesn't have a quick power function (math is all floaty). Since there are at most
	// 10 matches we're going to be really sloppy and hardcode.
	if matches == 0 {
		return 0
	}

	return 1 << (matches - 1)
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

	fmt.Printf("\nSolving\n%s\n", line)
	fmt.Printf("win: %v\n my: %v\n", winNumbers, myNumbers)
	fmt.Println(parts[0], "num matches:", matches)

	if matches > 10 {
		return 0, fmt.Errorf("too many matches. Had %d", matches)
	}

	return matches, nil

}
