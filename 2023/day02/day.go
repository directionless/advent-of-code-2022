package day02

import (
	"fmt"
	"strings"
)

type dayHandler struct {
	bag             pixelTyp
	possibleGameSum int
	totalPower      int
}

func New(bag pixelTyp) *dayHandler {
	h := &dayHandler{
		bag: bag,
	}

	return h
}

func (h *dayHandler) Consume(line []byte) error {
	if len(line) == 0 {
		return nil
	}

	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	chunks := strings.Split(string(line), ":")
	if len(chunks) != 2 {
		return fmt.Errorf(`malformed line "%s" didn't get 2 chunks`, line)
	}

	var gameNum int
	if n, err := fmt.Sscanf(chunks[0], "Game %d", &gameNum); err != nil {
		return fmt.Errorf(`parsing line "%s": %w`, line, err)
	} else if n != 1 {
		return fmt.Errorf(`malformed line "%s" matched %d times: %w`, line, n, err)
	}

	fmt.Println("found game", gameNum, ": ", chunks[1])

	minCubesForGame := pixelTyp{}

	possible := true
	games := strings.Split(chunks[1], ";")
	for _, g := range games {
		p, err := PixelFromString(g)
		if err != nil {
			return fmt.Errorf("parsing pixel from `%s` (game %d): %w", g, gameNum, err)
		}

		// part 1
		if !h.bag.Contains(p) {
			possible = false
		}

		// part 2
		if p.Red() > minCubesForGame.red {
			minCubesForGame.red = p.Red()
		}
		if p.Green() > minCubesForGame.green {
			minCubesForGame.green = p.Green()
		}
		if p.Blue() > minCubesForGame.blue {
			minCubesForGame.blue = p.Blue()
		}
	}

	// part 1
	if possible {
		fmt.Printf("Game %d is possible\n", gameNum)
		h.possibleGameSum += gameNum
	}

	// part 2
	power := minCubesForGame.Power()
	fmt.Printf("Game %d: %v Power %d\n", gameNum, minCubesForGame, power)
	h.totalPower += power

	return nil
}

func (h *dayHandler) AnswerPart1() any {
	return h.possibleGameSum
}

func (h *dayHandler) AnswerPart2() any {
	return h.totalPower
}

func (h *dayHandler) Print() {
	fmt.Printf("Part1: ???: %d\n", h.AnswerPart1())
	fmt.Printf("Part2: ???: %d\n", h.AnswerPart2())
}
