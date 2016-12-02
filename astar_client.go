package astar

// astar_clioent.go implements implements Pather for
// the sake of testing.  This functionality forms the back end for
// astart_client_test.go, and serves as an example for how to use A* for a graph.

// Nodes are called 'Nodes' and they have X, Y coordinates
// Edges are called 'Edges', they connect Nodes, and they have a cost
//
// The key differences between this example and the Tile world:
// 1) There is no grid.  Nodes have arbitrary coordinates.
// 2) Edges are not implied by the grid positions.  Instead edges are explicitly
//    modelled as 'Edges'.
//
// The key similarities between this example and the Tile world:
// 1) They both use Manhattan distance as their heuristic
// 2) Both implement Pather

// GobotWorld will eventually hold a map of type Node
type GobotWorld struct {
	//	nodes map[int]*Node		// not yet used
}

// Edge type connects two Nodes with a cost
type Edge struct {
	from *Node
	to   *Node
	Cost float64
}

// A Node is a place in a grid which implements Grapher.
type Node struct {

	// X and Y are the coordinates of the truck.
	X, Y int

	// array of tubes going to other trucks
	outTo []Edge

	label string
}

// PathNeighbors returns the neighbors of the Truck
func (t *Node) PathNeighbors() []Pather {

	neighbors := []Pather{}

	for _, edgeElement := range t.outTo {
		neighbors = append(neighbors, Pather(edgeElement.to))
	}
	return neighbors
}

// PathNeighborCost returns the cost of the edge leading to Node.
func (t *Node) PathNeighborCost(to Pather) float64 {

	for _, edgeElement := range (t).outTo {
		if Pather((edgeElement.to)) == to {
			return edgeElement.Cost
		}
	}
	return 10000000
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *Node) PathEstimatedCost(to Pather) float64 {

	toT := to.(*Node)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	r := float64(absX + absY)

	return r
}

// RenderPath renders a path on top of a Goreland world.
func (w GobotWorld) RenderPath(path []Pather) string {

	s := ""
	for _, p := range path {
		pT := p.(*Node)
		if pT.label != "END" {
			s = pT.label + "->" + s
		} else {
			s = s + pT.label
		}
	}
	return s
}
