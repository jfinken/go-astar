package astar

import (
	"testing"
)

func AddNode(x int, y int, label string) *Node {
	t1 := new(Node)
	t1.X = x
	t1.Y = y
	t1.label = label
	return t1
}

func AddEdge(t1, t2 *Node, cost float64) *Edge {
	edge1 := new(Edge)
	edge1.Cost = cost
	edge1.from = t1
	edge1.to = t2

	t1.outTo = append(t1.outTo, *edge1)

	return edge1
}

// Consider a world with Nodes and Edges, Edges each having a cost
//
//		   	 E
//			 |
//		   	 N9
//		  /	 |
//    	 /	 |
//       N7	 N8
//       |	 |
//N2--N1 N5--N6
//    |/ 	 |
//    S--N3--N4
//
// S=Start at (1,1)
// E=End at (3,5)
//
// N1 = (1, 2)
// N2 = (0, 2)
// N3 = (2, 1)
// N4 = (3, 1)
// N5 = (2, 2)
// N6 = (3, 2)
// N7 = (2, 3)
// N8 = (3, 3)
// N9 = (3, 4)
//
// S-M and M-E are clean clear edges. cost: 1
//
// S-E is either:
//
// 1) TestGraphPath_ShortDiagonal : diagonal cost is low at: 1.2
//    Solver should traverse diagonally through N5 and N7
//    Expect solution: Start, N5, N9, N9 End  Total cost: 4.2
//
// 1) TestGraphPath_LongDiagonal : diagonal is cost is very high
//    Solver should avoid those diagonal edges.
//    Expect solution total cost: 6.0

func createWorldGraphPathDiagonal(t *testing.T, diagonalCost float64, expectedDist float64) {

	world := new(GobotWorld)

	nStart := AddNode(0, 0, "START")
	n1 := AddNode(1, 2, "n1")
	n2 := AddNode(0, 2, "n2")
	n3 := AddNode(2, 1, "n3")
	n4 := AddNode(3, 1, "n4")
	n5 := AddNode(2, 2, "n5")
	n6 := AddNode(3, 2, "n6")
	n7 := AddNode(2, 3, "n7")
	n8 := AddNode(3, 3, "n8")
	n9 := AddNode(3, 4, "n9")
	nEnd := AddNode(1, 1, "END")

	AddEdge(nStart, n1, 1)
	AddEdge(nStart, n5, diagonalCost)
	AddEdge(nStart, n3, 1)
	AddEdge(n3, n4, 1)
	AddEdge(n1, n2, 1)
	AddEdge(n5, n6, 1)
	AddEdge(n5, n7, 1)
	AddEdge(n4, n6, 1)
	AddEdge(n6, n8, 1)
	AddEdge(n8, n9, 1)
	AddEdge(n7, n9, diagonalCost)
	AddEdge(n9, nEnd, 1)

	t.Logf("World.  Diagonal cost: %v\n\n", diagonalCost)

	p, dist, found := Path(nStart, nEnd)

	if !found {
		t.Log("Could not find a path")
	} else {
		t.Logf("Resulting path\n%s", world.RenderPath(p))
	}
	if !found && expectedDist >= 0 {
		t.Fatal("Could not find a path")
	}
	if found && dist != expectedDist {
		t.Fatalf("Expected dist to be %v but got %v", expectedDist, dist)
	}
}

func TestGraphPaths_ShortDiagonal(t *testing.T) {
	createWorldGraphPathDiagonal(t, 1.2, 4.4)
}
func TestGraphPaths_LongDiagonal(t *testing.T) {
	createWorldGraphPathDiagonal(t, 10000, 6.0)
}
