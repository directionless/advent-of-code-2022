package day15

import (
	"fmt"
	"strings"
)

type sensorType struct {
	Location location
	Beacon   location //nearest
	Distance int      // manhatten distance
}

func (s sensorType) String() string {
	return fmt.Sprintf("%s:%d", s.Location, s.Distance)
}

// RowY returns this sensors coverage on a given Y line. (-x, +x, ok)
func (s sensorType) RowY(y int) (int, int, bool) {
	if y > s.Location.Y+s.Distance || y < s.Location.Y-s.Distance {
		return 0, 0, false
	}

	// How much distance do we use up getting to Y (This will tell us how much remains for X)
	d := Abs(y-s.Location.Y) - s.Distance

	x1 := s.Location.X - d
	x2 := s.Location.X + d

	//fmt.Printf("sensor %s: row: %d, distance: %d\n", s, y, d)

	if x1 < x2 {
		return x1, x2, true
	}
	return x2, x1, true
}

type network struct {
	sensors []sensorType
}

func NewNetwork() *network {
	network := &network{
		sensors: make([]sensorType, 0),
	}

	return network
}

func (n *network) AddDetection(sensor, beacon location) {
	distance := manhattenDistance(sensor, beacon)

	n.sensors = append(
		n.sensors,
		sensorType{
			Location: sensor,
			Beacon:   beacon,
			Distance: distance,
		})
}

func (n *network) GetRowY(y int) (int, string) {
	row := map[int]rune{}
	var rowMin, rowMax int

	//beaconsThatMatter := make([]sensorType, 0)
	for _, s := range n.sensors {
		xMin, xMax, ok := s.RowY(y)
		if !ok {
			continue
		}

		fmt.Printf("sensor %s: %d - %d\n", s, xMin, xMax)

		if xMin < rowMin {
			rowMin = xMin
		}
		if xMax > rowMax {
			rowMax = xMax
		}

		for x := xMin; x <= xMax; x++ {
			row[x] = '#'
		}
	}

	// Add in existing beacons and sensors
	for _, s := range n.sensors {
		if s.Location.Y == y {
			row[s.Location.X] = 'S'
		}

		if s.Beacon.Y == y {
			row[s.Beacon.X] = 'B'
		}
	}

	// Count the spaces covered by detection, and make a pretty print too.
	var covered int
	var sb strings.Builder
	fmt.Fprintf(&sb, "row %d: ", y)
	for x := rowMin; x <= rowMax; x++ {
		if c, ok := row[x]; ok {
			if c == '#' {
				covered++
			}
			fmt.Fprintf(&sb, "%c", c)
		} else {
			fmt.Fprintf(&sb, ".")
		}
	}

	return covered, sb.String()
}
