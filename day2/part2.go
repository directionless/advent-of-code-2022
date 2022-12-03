package day2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func CalculateScorePart2(rd io.Reader) (int, error) {
	score := 0

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		// Split into two strings.
		line := strings.SplitN(scanner.Text(), " ", 2)
		if len(line) != 2 {
			continue
		}

		score += scoreFromRoundPart2(line[0], line[1])
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to scan: %w", err)
	}

	return score, nil
}

func scoreFromResult(res string) int {
	switch res {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	default:
		panic("unknown")
	}
}

// The Elf finishes helping with the tent and sneaks back over to you. "Anyway, the second column says how the round needs to end: X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

// It would be more clever to do this as an array, and then increment/decrement the index. But I'm going to hardcode a big case statement instead.
// var hands = []string{"A", "B", "C"}
func scoreFromRoundPart2(them, result string) int {
	score := scoreFromResult(result)
	switch {
	// lose
	case result == "X" && them == "A":
		score += shapeScore("C")
	case result == "X" && them == "B":
		score += shapeScore("A")
	case result == "X" && them == "C":
		score += shapeScore("B")

	// draw
	case result == "Y":
		score += shapeScore(them)

	// win
	case result == "Z" && them == "A":
		score += shapeScore("B")
	case result == "Z" && them == "B":
		score += shapeScore("C")
	case result == "Z" && them == "C":
		score += shapeScore("A")
	}

	return score
}
