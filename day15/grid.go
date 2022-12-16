package day15

import (
	"fmt"
	"strings"
)

type grid struct {
	detections map[location]location
	covered    map[location]bool
	pretty     map[location]rune

	minX int
	maxX int
	minY int
	maxY int
}

func NewGrid() *grid {
	g := &grid{
		detections: make(map[location]location),
		covered:    make(map[location]bool),
		pretty:     make(map[location]rune),
	}
	return g
}

func (g *grid) AddSensor(sensor, beacon location) {
	xdistance := beacon.X - sensor.X
	ydistance := beacon.Y - sensor.Y
	distance := Abs(xdistance) + Abs(ydistance)

	g.setBoundaries(beacon)
	g.setBoundaries(sensor)

	g.pretty[beacon] = 'B'
	g.pretty[sensor] = 'S'

	g.detections[beacon] = sensor

	// TODO
	for _, loc := range allWithinTaxiRange(sensor, distance) {
		g.covered[loc] = true

		if _, ok := g.pretty[loc]; !ok {
			g.pretty[loc] = '#'
		}

		// This is called way too many times. But :shrug:
		g.setBoundaries(loc)
	}
}

func (g *grid) CoveredInRow(y int) int {
	count := 0
	loc := location{0, y}
	for x := g.minX; x <= g.maxX; x++ {
		loc.X = x
		fmt.Printf("%c", g.pretty[loc])
		// Need to differenciate betweened covered and existing beacons
		if g.pretty[loc] == '#' { // g.covered[loc] {
			count++
		}
	}
	return count
}

func allWithinTaxiRange(loc location, totalDistance int) []location {
	var locations []location

	for d := 1; d <= totalDistance; d++ {
		for dx := 0; dx <= d; dx++ {
			dy := d - dx
			locations = append(locations, location{X: loc.X + dx, Y: loc.Y + dy})
			locations = append(locations, location{X: loc.X + dx, Y: loc.Y - dy})
			locations = append(locations, location{X: loc.X - dx, Y: loc.Y + dy})
			locations = append(locations, location{X: loc.X - dx, Y: loc.Y - dy})
		}
	}

	return locations
}

func (g *grid) setBoundaries(loc location) {
	if loc.X < g.minX {
		g.minX = loc.X
	}
	if loc.X > g.maxX {
		g.maxX = loc.X
	}
	if loc.Y < g.minY {
		g.minY = loc.Y
	}
	if loc.Y > g.maxY {
		g.maxY = loc.Y
	}
}

func (g *grid) String() string {
	var sb strings.Builder

	for h := 0; h < 2; h++ {
		fmt.Fprintf(&sb, "      ")
		for x := g.minX; x <= g.maxX; x++ {
			switch h {
			case 0:
				if x%10 == 0 {
					fmt.Fprintf(&sb, "%-10d", x)
				}
			case 1:
				fmt.Fprintf(&sb, "-")
			case 2:
				fmt.Fprintf(&sb, " ")
			}
		}
		fmt.Fprintf(&sb, "\n")
	}

	for y := g.minY; y <= g.maxY; y++ {
		fmt.Fprintf(&sb, "%3d | ", y)
		for x := g.minX; x <= g.maxX; x++ {
			loc := location{x, y}
			if v, ok := g.pretty[loc]; ok {
				fmt.Fprintf(&sb, "%c", v)
			} else {
				fmt.Fprintf(&sb, ".")
			}
		}
		fmt.Fprintf(&sb, "\n")
	}
	return sb.String()
}
