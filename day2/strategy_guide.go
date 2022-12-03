package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.
//
// The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious, so the responses must have been carefully chosen.
//
// The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
func scoreFromRound(them, me string) int {
	var score int

	// Determine winner. Don't need to be too fancy. Just use switch.
	// Rock: A, X
	// Paper: B, Y
	// Scissors: C, Z
	switch {
	case them == "A" && me == "X":
		score += 3
	case them == "A" && me == "Y":
		score += 6
	case them == "A" && me == "Z":
		score += 0

	case them == "B" && me == "X":
		score += 0
	case them == "B" && me == "Y":
		score += 3
	case them == "B" && me == "Z":
		score += 6

	case them == "C" && me == "X":
		score += 6
	case them == "C" && me == "Y":
		score += 0
	case them == "C" && me == "Z":
		score += 3
	}

	score += shapeScore(me)

	return score
}

func shapeScore(shape string) int {
	switch shape {
	case "X": // rock
		return 1
	case "Y": // paper
		return 2
	case "Z": // scissors
		return 3
	default:
		panic("unknown")
	}
}

func calculateScore(rd io.Reader) (int, error) {
	score := 0

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		// Split into two strings.
		line := strings.SplitN(scanner.Text(), " ", 2)
		if len(line) != 2 {
			continue
		}

		score += scoreFromRound(line[0], line[1])
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to scan: %w", err)
	}

	return score, nil

}
