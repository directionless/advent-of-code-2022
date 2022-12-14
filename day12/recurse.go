package day12

import (
	"container/list"
	"fmt"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
)

// (really this is part of day handler)

type location struct {
	X               int
	Y               int
	H               byte
	nodeID          int64
	DistanceToStart int
}

// ID is compatible with the gonum/graph pkg
func (l location) ID() int64 {
	return l.nodeID
}

func (l location) String() string {
	return fmt.Sprintf("%d,%d,%c", l.X, l.Y, l.H)
}

// This is kinda silly. The `location` is used, not as a graph node, but as a simple encapsulation of
// the X,Y pair. However, the `*location` pointer _is_ a graph node.
type grid struct {
	start  *location
	end    *location
	grid   map[location]*location
	width  int
	height int

	// gonum/graph stuff
	nodeNum int64
	graph   *simple.DirectedGraph

	// handrolled BFS
	bfsQueue  *list.List
	visited   map[*location]bool
	paths     map[*location][]*location
	bfsSlice  []*location
	bfsSliceI int
}

func NewGrid() *grid {
	g := &grid{
		grid:    make(map[location]*location),
		graph:   simple.NewDirectedGraph(),
		nodeNum: 100,
	}
	return g
}

func (g *grid) Node(x, y int) *location {
	return g.grid[location{X: x, Y: y}]
}

func (g *grid) AddRow(row []byte) error {
	if g.width == 0 {
		g.width = len(row)
	} else {
		if g.width != len(row) {
			return fmt.Errorf("line length mismatch: %d != %d", g.width, len(row))
		}
	}

	for x, c := range row {
		g.nodeNum += 1
		loc := &location{
			X:      x,
			Y:      g.height,
			H:      c,
			nodeID: g.nodeNum,
		}
		g.graph.AddNode(loc)
		g.grid[location{X: x, Y: g.height}] = loc

		switch c {
		case 'S':
			g.start = loc
		case 'E':
			g.end = loc
		}
	}
	g.height++
	return nil
}

// BuildGraph builds the gonum/graph object for usage.
func (g *grid) BuildGraph() error {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			cur := g.grid[location{X: x, Y: y}]
			if cur == nil {
				return fmt.Errorf("no node at %d,%d", x, y)
			}

			for _, neigh := range g.Neighbors(cur) {
				g.graph.SetEdge(simple.Edge{cur, neigh})
			}
		}
	}

	return nil
}

// Neighbors returns the reachable neighbors of a given node.
func (g *grid) Neighbors(cur *location) []*location {
	possibilities := []location{
		{X: cur.X, Y: cur.Y - 1},
		{X: cur.X, Y: cur.Y + 1},
		{X: cur.X - 1, Y: cur.Y},
		{X: cur.X + 1, Y: cur.Y},
	}

	var neighbors []*location

	for _, loc := range possibilities {
		node, ok := g.grid[loc]
		if !ok {
			continue
		}

		myHeight := cur.H
		targetH := node.H

		//fmt.Printf("examining %c with %c\n", cur.H, node.H)

		// S is like 'a', 'E' is like 'z'
		if myHeight == 'S' {
			myHeight = 'a'
		}
		if targetH == 'S' {
			targetH = 'a'
		}

		if myHeight == 'E' {
			myHeight = 'z'
		}
		if targetH == 'E' {
			targetH = 'z'
		}

		// We can step, at most, one up
		if int(myHeight)-1 <= int(targetH) {
			neighbors = append(neighbors, node)
			fmt.Printf("Found neighbor %s from %s. height: %c > %c \n", node, cur, myHeight, node.H)
			continue
		}

	}

	return neighbors
}

func (g *grid) String() string {
	var sb strings.Builder

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			fmt.Fprintf(&sb, "%s", string(g.grid[location{X: x, Y: y}].H))
		}
		fmt.Fprintf(&sb, "\n")
	}

	return sb.String()
}

func (g *grid) ResetBFS() {
	g.visited = make(map[*location]bool, len(g.grid))
	g.bfsQueue = list.New()
	g.paths = make(map[*location][]*location, len(g.grid))

	g.start.DistanceToStart = 0
	g.visited[g.end] = true
	g.bfsQueue.PushBack(g.end)
}

func (g *grid) BFS(lowNotStart bool) int {
	for g.bfsQueue.Len() > 0 {
		// Note the lack of nil check here. live dangereously
		cur := g.bfsQueue.Front().Value.(*location)
		g.bfsQueue.Remove(g.bfsQueue.Front())

		fmt.Printf("BFS check %s\n", cur)

		if !lowNotStart && cur.X == g.start.X && cur.Y == g.start.Y {
			fmt.Printf("found: %s. It's %s\n", g.end, cur)
			g.PrintPath(cur)
			return cur.DistanceToStart
		}

		if lowNotStart && cur.H == 'a' {
			fmt.Printf("found: %s. It's %s\n", g.end, cur)
			g.PrintPath(cur)
			return cur.DistanceToStart

		}

		neighbors := g.Neighbors(cur)

		for _, n := range neighbors {
			if g.visited[n] {
				continue
			}
			n.DistanceToStart = cur.DistanceToStart + 1
			g.paths[n] = append(g.paths[cur], n)
			g.bfsQueue.PushBack(n)
			g.visited[n] = true
		}
	}
	return -1
}

func (g *grid) PrintPath(l *location) {
	out := map[[2]int]byte{}
	for _, p := range g.paths[l] {
		fmt.Printf("%s\n", p)
		out[[2]int{p.X, p.Y}] = p.H
	}

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if out[[2]int{x, y}] != 0 {
				fmt.Printf("%c", out[[2]int{x, y}])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
