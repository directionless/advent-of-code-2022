package crt

import (
	"fmt"
	"io"
	"os"
)

type pixel struct {
	color bool
}

func (p *pixel) Light() {
	p.color = true
}

func (p pixel) Output() string {
	if p.color {
		return "#"
	}
	return "."
}

type CrtStruct struct {
	width  int
	height int
	grid   map[[2]int]*pixel
}

func New(width, height int) *CrtStruct {
	crt := &CrtStruct{
		width:  width,
		height: height,
		grid:   make(map[[2]int]*pixel, width*height),
	}

	for y := crt.height - 1; y >= 0; y-- {
		for x := 0; x < crt.width; x++ {
			crt.grid[[2]int{x, y}] = &pixel{}
		}
	}

	return crt
}

func (crt CrtStruct) OutputF(f io.Writer) {
	//for y := crt.height - 1; y >= 0; y-- {
	for y := 0; y < crt.height; y++ {

		for x := 0; x < crt.width; x++ {
			loc := [2]int{x, y}
			p := crt.grid[loc]
			if p == nil {
				fmt.Fprint(f, " ")
			} else {
				fmt.Fprintf(f, p.Output())
			}
		}

		fmt.Fprintf(f, "\n")
	}
}

func (crt CrtStruct) Output() {
	for x := 0; x < crt.width; x++ {
		fmt.Print("=")
	}
	fmt.Println()

	crt.OutputF(os.Stdout)

	for x := 0; x < crt.width; x++ {
		fmt.Print("=")
	}
	fmt.Println()
}

// Examine allows us to tie the CRT to the CPU. It is a callback from the CPU. Drawing
// rules as follow:
//
// 1. the CRT draws a single pixel during each cycle
// 2. AFAICT nothing turns them off
// 3. the sprite is 3 pixels wide
// 4. the X register sets the horizontal position of the middle of that sprite
func (crt *CrtStruct) Examine(cycle int, x int) {
	pos := posFromCycle(cycle)

	// Does pos overlap with the sprite? Sprite is 3 pixels wide, set by X
	if pos[0] >= x-1 && pos[0] <= x+1 {
		fmt.Printf("cycle %d, x %d: lighting for pixel at %v\n", cycle, x, pos[0])
		if crt.grid[pos] != nil {
			crt.grid[pos].Light()
		} else {
			// uh oh
			fmt.Print(" ")
		}
	} else {
		fmt.Printf("cycle %d, x %d: no light\n", cycle, x)
	}
}

func posFromCycle(cycle int) [2]int {
	// 0 vs 1 index shenangins
	cycle = cycle - 1

	x := cycle % 40
	y := cycle / 40
	return [2]int{x, y}
}
