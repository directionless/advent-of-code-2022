package day16

import (
	"fmt"
	"math"
	"math/bits"

	graphpath "gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/stat/combin"
)

// Thank you:
// - https://gist.github.com/ynyBonfennil/7c6a6b7cae7727efea274692643281ac

type valveType struct {
	Name   string
	Rate   int
	nodeID int64
}

// ID is compatible with the gonum/graph pkg
func (v valveType) ID() int64 {
	return v.nodeID
}

func (v valveType) String() string {
	return v.Name
}

const (
	maxTime = 30
)

type pressureFromPath struct {
	T    int // time taken
	MP   int // max pressure
	Path string
}

type valvePath string

type network struct {
	graph          *simple.UndirectedGraph
	valuableGraph  *simple.WeightedUndirectedGraph
	nodes          map[string]*valveType // name to valves
	valuable       []*valveType
	costs          map[*valveType]map[*valveType]int
	bitlen         int
	pressureCache  map[int64]pressureFromPath
	pressureCache2 map[valvePath]pressureFromPath
}

func newNetwork() *network {
	return &network{
		graph:          simple.NewUndirectedGraph(),
		valuableGraph:  simple.NewWeightedUndirectedGraph(0, math.Inf(1)),
		nodes:          make(map[string]*valveType),
		valuable:       make([]*valveType, 0),
		costs:          make(map[*valveType]map[*valveType]int),
		pressureCache:  make(map[int64]pressureFromPath, 0),
		pressureCache2: make(map[valvePath]pressureFromPath, 0),
	}
}

func (n *network) AddValve(name string, rate int, neighbors [][]byte) {
	v := n.findOrCreateValve(name)
	v.Rate = rate

	if v.Rate > 0 || v.Name == "AA" {
		n.valuable = append(n.valuable, v)
		n.valuableGraph.AddNode(v)
	}

	for _, neighbor := range neighbors {
		n.AddEdge(v, n.findOrCreateValve(string(neighbor)))
	}
}

func (n *network) AddEdge(a, b *valveType) {
	n.graph.SetEdge(simple.Edge{F: a, T: b})
}

func (n *network) findOrCreateValve(name string) *valveType {
	v, ok := n.nodes[name]
	if ok {
		return v
	}

	v = &valveType{
		Name:   name,
		nodeID: int64(len(n.nodes)),
	}

	if l := bits.Len(uint(v.ID())); l > n.bitlen {
		n.bitlen = l
	}

	n.nodes[name] = v

	return v
}

// pathToBits converts a path of nodes, to a bit string, this is used as a map key.
// The converstion treats each place in the path as the amount to left shift.
func (n *network) pathToBits(path []*valveType) int64 {
	var hashkey int64
	for i, v := range path {
		hashkey = hashkey | v.ID()<<i
	}
	return hashkey
}

func (n *network) MaxPressure(path []*valveType) pressureFromPath {
	if pt, ok := n.pressureCache[n.pathToBits(path)]; ok {
		//fmt.Println("Found pressure in cache")
		return pt
	}

	if len(path) == 0 || path == nil {
		return pressureFromPath{}
	}

	// If this is the first node, in our path, it may or may not be worth opening.
	// (AA is special). Return appropriate values.
	if len(path) == 1 {
		if path[0].Rate > 0 {
			// If there's a valve here, open it.
			return pressureFromPath{
				T:    1,
				MP:   path[0].Rate,
				Path: fmt.Sprintf("%s", path[0]),
			}
		} else {
			return pressureFromPath{}
		}
	}

	// If we get here, then we have enough to look at the prior pressure
	// and node, and whatnot
	pNode := path[len(path)-2]
	pt := n.MaxPressure(path[0 : len(path)-1])

	travelCost := n.costs[pNode][path[len(path)-1]]
	pt.T += travelCost

	// Add one to open the valve
	pt.T += 1

	// add path
	pt.Path += fmt.Sprintf("%s", pNode)

	// Check time budget. If we have time, open the valve and calculate the MP it contributes
	if pt.T < maxTime {
		ticks := maxTime - travelCost - 1
		pt.MP += pNode.Rate * ticks
	} else {
		pt.MP = -1
	}

	n.pressureCache[n.pathToBits(path)] = pt

	return pt
}

func (n *network) SolvePart1() (pressureFromPath, error) {
	if err := n.CaclculateCosts(); err != nil {
		return pressureFromPath{}, fmt.Errorf("calculating costs: %w", err)
	}

	// We know the cost between each valuable node, and the flow rate. Now we need
	// to test all possible paths through the graph. We cannot simply test all
	// permutations. At 15ish valuable notes, it is too many. But, there are
	// more clever patterns we can take. Remember, we must start a AA, so it's a
	// litle smaller.
	needed := combin.NumPermutations(len(n.valuable), len(n.valuable))
	fmt.Printf("We're going to need to test %d permutations\n", needed)

	// Going to need to roll a DFS :<
	//
	// remaining would be cleaner as a bitfield, but my bits are failing me.
	// So ugly string garbage.
	startPath := []*valveType{n.nodes["AA"]}
	remainingNodes := 0 | (1 << uint(n.nodes["AA"].ID()))
	if len(n.valuable) > 128 {
		return pressureFromPath{}, fmt.Errorf("too many valuable nodes. Think harder")
	}
	n.dfs(startPath, remainingNodes)

	fmt.Printf("Cache size: %d\n", len(n.pressureCache))

	//find biggest in cache
	var max pressureFromPath
	for _, pt := range n.pressureCache {
		if pt.MP > max.MP {
			max = pt
		}
	}

	return max, nil
}

func (n *network) dfs(path []*valveType, remaining int) {
	//fmt.Printf("dfs %v. Remaning %b\n", path, remaining)
	pt := n.MaxPressure(path)

	// early return if no good
	if pt.MP < 0 {
		//fmt.Println("MP negative. Give up!")
		return
	}

	for _, v := range n.valuable {
		if remaining&(1<<uint(v.ID())) != 0 {
			// this isn'r working right. So ugly parsing for now
			continue
		}

		//fmt.Printf("Trying node %s\n", v)
		newPath := append(path, v)
		newRemaining := remaining | (1 << uint(v.ID()))
		n.dfs(newPath, newRemaining)
	}
}

func (n *network) CaclculateCosts() error {
	allShortest, ok := graphpath.FloydWarshall(n.graph)
	if !ok {
		return fmt.Errorf("failed to calculate shortest paths")
	}

	for _, v1 := range n.valuable {
		n.costs[v1] = make(map[*valveType]int, len(n.valuable))
		for _, v2 := range n.valuable {
			if v1 == v2 {
				continue
			}

			// The path returned here includes the head and tail, the weight appear
			// to be the number of edges. Since we want the number of edges, we use weight
			_, weight, _ := allShortest.Between(v1.ID(), v2.ID())
			fmt.Printf("From %s to %s: weight %d\n", v1, v2, int(weight))
			n.costs[v1][v2] = int(weight)
			weightedEdge := simple.WeightedEdge{F: v1, T: v2, W: weight}
			n.valuableGraph.SetWeightedEdge(weightedEdge)
		}
	}

	return nil
}
