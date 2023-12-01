package day17

import (
	"fmt"
	"strings"
)

const (
	chamberWidth = 7
)

type pieceShape int

const (
	horizontalPiece pieceShape = iota
	plusPiece
	lPiece
	verticalPiece
	squarePiece
)

func (p pieceShape) String() string {
	switch p {
	case horizontalPiece:
		return "-"
	case plusPiece:
		return "+"
	case lPiece:
		return "L"
	case verticalPiece:
		return "|"
	case squarePiece:
		return "#"
	}
	return "?"
}

func (p pieceShape) Shape() [4][4]bool {
	// These are inverted, so that the first row, is the "bottom"
	switch p {
	case horizontalPiece:
		return [4][4]bool{
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}
	case plusPiece:
		return [4][4]bool{
			{false, true, false, false},
			{true, true, true, false},
			{false, true, false, false},
			{false, false, false, false},
		}
	case lPiece:
		return [4][4]bool{
			{false, false, false, false},
			{true, true, true, false},
			{false, false, true, false},
			{false, false, true, false},
		}
	case verticalPiece:
		return [4][4]bool{
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
		}
	case squarePiece:
		return [4][4]bool{
			{true, true, false, false},
			{true, true, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}
	}
	return [4][4]bool{}
}

type chamber struct {
	jetpattern []byte
	jetIdx     int
	stack      [][chamberWidth]bool
	startY     int
}

// AddPiece adds a piece according to the rules.
//
// The tall, vertical chamber is exactly seven units wide. Each rock
// appears so that its left edge is two units away from the left wall
// and its bottom edge is three units above the highest rock in the
// room (or the floor, if there isn't one).
//
// After a rock appears, it alternates between being pushed by a jet of
// hot gas one unit (in the direction indicated by the next symbol in the
// jet pattern) and then falling one unit down.
func (c *chamber) AddPiece(p pieceShape) {
	if c.startY < 0 {
		c.startY = 2
	}

	startX := 2

	shape := p.Shape()

}

func (c *chamber) CheckCollisions(x, y int, shape [4][4]bool) (wallCollide, floorColide bool) {
	for dy := 0; dy < 4; dy++ {
		for dx := 0; dx < 4; dx++ {

		}
	}
}

func (c *chamber) String() string {
	var sb strings.Builder
	for i := len(c.stack) - 1; i >= 0; i-- {
		fmt.Fprintf(sb, "%s\n", c.stack[i])
	}
	return sb.String()
}
