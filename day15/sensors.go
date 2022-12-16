package day15

import (
	"fmt"
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
	return x1, x2, true
}

// GetHole examines the outer edges (previously calculated) to find one that is
// not contained within a sensor circle
func (n *network) GetHole() (location, error) {
	fmt.Printf("Starting to look for holes\n")
	fmt.Printf("There are %d to examine\n", len(n.outerPoints))

	count := len(n.outerPoints)

	found := make([]location, 0)

	for loc, c := range n.outerPoints {
		count--

		if count%1000 == 0 {
			fmt.Printf("   %d left to examine\n", count)
		}

		// If we assume our mystery hole isn't on a terminal edge, then it will inherently be in
		// more than one sensors outer ring. So we can early return if we only match one.
		if c == 1 {
			continue
		}

		if ok := n.sensorsContain(loc); ok {
			continue
		}

		found = append(found, loc)
		fmt.Printf("found: %s\n", loc)
	}

	if len(found) != 1 {
		fmt.Printf("Found too many holes. Got %d\n", len(found))
		return location{}, fmt.Errorf("wrong number of holds: Got %d", len(found))
	}
	return found[0], nil
}

func (n *network) sensorsContain(loc location) bool {
	for _, s := range n.sensors {
		x1, x2, ok := s.RowY(loc.Y)
		if !ok {
			continue
		}

		//fmt.Printf("loc is %s, sensor is %d - %d\n", loc, x1, x2)

		if x1 <= loc.X && loc.X <= x2 {
			return true
		}

		if x2 <= loc.X && loc.X <= x1 {
			return true
		}

	}

	return false
}

type outerPointHolder interface {
	AddOuterPoint(l location)
}

// FindEdges finds the outer edges of the sensor coverage. This is defines using the
// manhatten distance from the sensor to the nearest beacon.
func (s sensorType) FindEdges(oph outerPointHolder) {

	for dy := 0; dy <= s.Distance; dy++ {
		dx := s.Distance - dy

		// With the distance in X and Y, we can now make edges
		//
		// With some weirdness since we want the _outer_ edges. The 1 here, on the X axis,
		// gets the x things. We need to special case the y outer

		x1 := s.Location.X - dx - 1
		x2 := s.Location.X + dx + 1

		y1 := s.Location.Y - dy
		y2 := s.Location.Y + dy

		oph.AddOuterPoint(location{x1, y1})
		oph.AddOuterPoint(location{x2, y1})
		oph.AddOuterPoint(location{x1, y2})
		oph.AddOuterPoint(location{x2, y2})
	}

	// Get the top and bottom outer
	oph.AddOuterPoint(location{s.Location.X, s.Location.Y + s.Distance + 1})
	oph.AddOuterPoint(location{s.Location.X, s.Location.Y - s.Distance - 1})

}

type network struct {
	sensors []sensorType

	// Holds the number of times we've seen each point. This will let us search for
	// the singular one.
	part2Max    int
	outerPoints map[location]int
}

func NewNetwork() *network {
	network := &network{
		sensors: make([]sensorType, 0),
	}

	return network
}

func (n *network) AddOuterPoint(l location) {
	//fmt.Printf("outer pointL %s\n", l)

	if l.Y < 0 || l.X < 0 {
		return
	}

	if l.Y > n.part2Max {
		return
	}
	if l.X > n.part2Max {
		return
	}

	n.outerPoints[l] += 1
}

func (n *network) FindHoles(part2Max int) (location, error) {
	n.part2Max = part2Max
	guessAllocation := n.part2Max * 10
	n.outerPoints = make(map[location]int, guessAllocation)

	fmt.Println("Initializing hole finder:")
	for _, s := range n.sensors {
		fmt.Printf("   %s...\n", s)
		s.FindEdges(n)
	}
	fmt.Println("Done initializing hole finder.")

	return n.GetHole()

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

func (n *network) GetRowY(y int) int {
	row := map[int]bool{}

	for _, s := range n.sensors {
		xMin, xMax, ok := s.RowY(y)
		if !ok {
			continue
		}

		//fmt.Printf("sensor %s: %d - %d\n", s, xMin, xMax)

		for x := xMin; x <= xMax; x++ {
			row[x] = true
		}
	}

	// Add in existing beacons and sensors
	for _, s := range n.sensors {
		if s.Location.Y == y {
			row[s.Location.X] = false
		}

		if s.Beacon.Y == y {
			row[s.Beacon.X] = false
		}
	}

	// Count the spaces covered by detection.
	var covered int
	for _, c := range row {
		if c {
			covered++
		}
	}

	return covered
}
