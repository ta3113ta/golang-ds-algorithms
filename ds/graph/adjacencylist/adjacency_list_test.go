package adjacencylist

import "testing"

func TestAddEdge(t *testing.T) {
	g := NewGraph[int]()

	g.AddEdge(0, 1)
	g.AddEdge(0, 4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	if g.AdjacencyList[0].Vertex != 4 {
		t.Error("Expected 4")
	}

	if g.AdjacencyList[1].Vertex != 3 {
		t.Error("Expected 3")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := NewGraph[int]()

	g.AddEdge(0, 1)
	g.AddEdge(0, 4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	g.RemoveEdge(1, 3)

	if g.AdjacencyList[1].Vertex != 2 {
		t.Error("Expected 2")
	}
}

func TestRemoveVertex(t *testing.T) {
	g := NewGraph[int]()

	g.AddEdge(0, 1)
	g.AddEdge(0, 4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	g.RemoveVertex(1)

	if g.AdjacencyList[0].Vertex != 4 {
		t.Error("Expected 4")
	}
}

func TestDFS(t *testing.T) {
	g := NewGraph[int]()

	g.AddEdge(0, 1)
	g.AddEdge(0, 4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	result := g.DFS(0)

	if len(result) != 5 {
		t.Error("Expected 5")
	}

	if result[0] != 0 {
		t.Error("Expected 0")
	}

	if result[1] != 4 {
		t.Error("Expected 4")
	}

	if result[2] != 1 {
		t.Error("Expected 1")
	}

	if result[3] != 3 {
		t.Error("Expected 3")
	}

	if result[4] != 2 {
		t.Error("Expected 2")
	}
}
