package day12

import (
	"fmt"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

// (really this is part of day handler)

type location struct {
	X      int
	Y      int
	H      byte
	nodeID int64
}

// ID is compatible with the graph library
func (l location) ID() int64 {
	return l.nodeID
}

func (l location) String() string {
	return fmt.Sprintf("%d,%d,%c", l.X, l.Y, l.H)
}

type path struct {
	start    location
	end      location
	distance int
	Steps    []location
}

// This is kinda silly. The `location` is used, not as a graph node, but as a simple encapsulation of
// the X,Y pair. However, the `*location` pointer _is_ a graph node.
type grid struct {
	start  *location
	end    *location
	grid   map[location]*location
	width  int
	height int

	nodeNum int64
	graph   *simple.DirectedGraph
}

func NewGrid() *grid {
	g := &grid{
		grid:    make(map[location]*location),
		graph:   simple.NewDirectedGraph(),
		nodeNum: 100,
	}
	return g
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

		// S is like 'a', 'E' is like 'z'
		if myHeight == 'S' {
			myHeight = 'a'
		}

		if myHeight == 'E' {
			myHeight = 'z'
		}

		// Stepping down is okay
		if int(myHeight) >= int(node.H) {
			neighbors = append(neighbors, node)
			continue
		}

		// We can step, at most, one up
		if int(myHeight)+1 >= int(node.H) {
			neighbors = append(neighbors, node)
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

func (g *grid) Solve() error {
	if err := g.BuildGraph(); err != nil {
		return err
	}

	// I don't understand what I would use for the traverse function. It seems to be
	// an additional constraint on the edges
	//traverseFn := func(_ graph.Edge) bool {
	//	return true
	//}

	//spew.Dump(g.graph.Edges())

	visitFn := func(n graph.Node) {
		fmt.Printf("visited %s\n", n)
	}

	bfs := traverse.BreadthFirst{
		//Traverse: traverseFn,
		Visit: visitFn,
	}

	var got [][]int64
	_ = bfs.Walk(
		g.graph,
		g.start,
		func(n graph.Node, d int) bool {
			if n == g.end {
				return true
			}

			if d >= len(got) {
				got = append(got, []int64(nil))
			}
			got[d] = append(got[d], n.ID())

			return false
		})

	fmt.Printf("got: %d\n", len(got))

	//spew.Dump(got)

	//found := graphpath.YenKShortestPaths(g.graph, 1, g.start, g.end)
	//spew.Dump(found)

	return nil
}

/*

func (g *grid) Solve() error {
	if err := g.BuildGraph(); err != nil {
		return err
	}

	if g.grid[g.start] != 'S' {
		return fmt.Errorf("start not set")
	}

	fmt.Printf("Starting at %s: %c\n", g.start, g.grid[g.start])

	paths := g.Recurse(g.start, path{})

	fmt.Printf("Found %d possible paths\n", len(paths))

	for _, p := range paths {
		fmt.Printf("Path is %d long\n", len(p.Steps))
		fmt.Printf("%v", p.Steps)
	}
	fmt.Printf("\n")
	return nil
}

func (g *grid) Recurse(cur location, p path) []path {
	if g.visted[cur] {
		return nil
	}

	g.visted[cur] = true

	possibilities := []location{
		{X: cur.X, Y: cur.Y - 1},
		{X: cur.X, Y: cur.Y + 1},
		{X: cur.X - 1, Y: cur.Y},
		{X: cur.X + 1, Y: cur.Y},
	}

	// Denote this step.
	p.Steps = append(p.Steps, cur)

	var paths []path

	for _, possibility := range possibilities {
		if possibility.X < 0 || possibility.Y < 0 {
			continue
		}

		if g.grid[possibility] == 0 {
			continue
		}

		// Have we been here before? This should prevent us from backtracking
		// This probably won't work. Consider a path that has a long loop and a
		// cross. We might end up on the loop, and we'll over count.
		if g.visted[possibility] {
			//continue
		}

		// Are we done?
		if g.grid[possibility] == 'E' {
			return []path{p}
		}

		// Do these points connect?
		if g.grid[possibility] != 'S' &&
			g.grid[possibility] != 'E' &&
			g.grid[cur] != 'S' &&
			g.grid[cur] != 'E' {

			terrainDiff := int(g.grid[possibility]) - int(g.grid[cur])
			if terrainDiff < -1 && terrainDiff > 1 {
				fmt.Printf("Cannot travel between %s(%c) and %s(%c)\n",
					cur, g.grid[cur], possibility, g.grid[possibility],
				)
				continue
			}
		}

		// Right now, just for debugging
		possibility.H = g.grid[possibility]

		// Recurse
		paths = append(paths, g.Recurse(possibility, p)...)
	}

	return paths
}
*/
