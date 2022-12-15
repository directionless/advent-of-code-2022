// package aocgrid is a simple grid implementation. It provides a little logic around
// cardinal directions, and a String function
package aocgrid

import (
	"fmt"
	"strings"
)

type gridValue interface {
	String() string
}

type Grid struct {
	minX       int
	maxX       int
	minY       int
	maxY       int
	havePoints bool
	grid       map[[2]int]gridValue
}

func New() *Grid {
	return &Grid{
		grid: make(map[[2]int]gridValue),
	}
}

func (g *Grid) AddRow() {
	panic("not implemented")
}

func (g *Grid) SetYMin(y int) {
	g.minY = y
}

func (g *Grid) Set(x, y int, v gridValue) {
	g.grid[[2]int{x, y}] = v

	if !g.havePoints {
		g.minX = x
		g.maxX = x
		g.minY = y
		g.maxY = y
		g.havePoints = true
		return
	}

	if x > g.maxX {
		g.maxX = x
	}
	if x < g.minX {
		g.minX = x
	}
	if y > g.maxY {
		g.maxY = y
	}
	if y < g.minY {
		g.minY = y
	}
}

func (g *Grid) FillRemaining(v gridValue) {
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if _, ok := g.grid[[2]int{x, y}]; !ok {
				g.grid[[2]int{x, y}] = v //g.Set(x, y, v)
			}
		}
	}
}

func (g *Grid) String() string {
	var sb strings.Builder

	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if v, ok := g.grid[[2]int{x, y}]; ok {
				fmt.Fprintf(&sb, "%s", v)
			} else {
				fmt.Fprintf(&sb, " ")
			}
		}
		fmt.Fprintf(&sb, "\n")
	}
	return sb.String()
}

func (g *Grid) SetStraightLine(x1, y1, x2, y2 int, v gridValue) error {
	switch {
	case x1 == x2 && y1 == y2:
		g.Set(x1, y1, v)
		return nil
	case x1 == x2:
		yS, yE := y1, y2
		if y1 > y2 {
			yS, yE = y2, y1
		}
		for y := yS; y <= yE; y++ {
			g.Set(x1, y, v)
		}
		return nil
	case y1 == y2:
		xS, xE := x1, x2
		if x1 > x2 {
			xS, xE = x2, x1
		}
		for x := xS; x <= xE; x++ {
			g.Set(x, y1, v)
		}
		return nil
	default:
		return fmt.Errorf("line not straight (%d,%d) -> (%d, %d)", x1, y1, x2, y2)
	}
}

func (g *Grid) GetCoordinates(x, y int) (int, int, bool) {
	if x > g.maxX || x < g.minX || y > g.maxY || y < g.minY {
		return 0, 0, false
	}

	if _, ok := g.grid[[2]int{x, y}]; !ok {
		return 0, 0, false
	}
	return x, y, true
}

func (g *Grid) Up(x, y int) (int, int, bool) {
	return g.GetCoordinates(x, y-1)
}

func (g *Grid) Down(x, y int) (int, int, bool) {
	return g.GetCoordinates(x, y+1)
}

func (g *Grid) DownLeft(x, y int) (int, int, bool) {
	return g.GetCoordinates(x-1, y+1)
}

func (g *Grid) DownRight(x, y int) (int, int, bool) {
	return g.GetCoordinates(x+1, y+1)
}

func (g *Grid) Left(x, y int) (int, int, bool) {
	return g.GetCoordinates(x-1, y)
}

func (g *Grid) Right(x, y int) (int, int, bool) {
	return g.GetCoordinates(x+1, y)
}

func (g *Grid) Look(x, y int) gridValue {
	if x > g.maxX || x < g.minX || y > g.maxY || y < g.minY {
		return nil
	}

	val, ok := g.grid[[2]int{x, y}]
	if !ok {
		return nil
	}
	return val
}

func (g *Grid) LookUp(x, y int) gridValue {
	return g.Look(x, y-1)
}

func (g *Grid) LookDown(x, y int) gridValue {
	return g.Look(x, y+1)
}

func (g *Grid) LookDownLeft(x, y int) gridValue {
	return g.Look(x-1, y+1)
}

func (g *Grid) LookDownRight(x, y int) gridValue {
	return g.Look(x+1, y+1)
}

func (g *Grid) LookLeft(x, y int) gridValue {
	return g.Look(x-1, y)
}

func (g *Grid) LookRight(x, y int) gridValue {
	return g.Look(x+1, y)
}
