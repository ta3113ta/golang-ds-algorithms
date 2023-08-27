package adjacencylist

type Node[T comparable] struct {
	Vertex T
	Next   *Node[T]
}

type Graph[T comparable] struct {
	AdjacencyList map[T]*Node[T]
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		AdjacencyList: make(map[T]*Node[T]),
	}
}

func (g *Graph[T]) AddEdge(src, dest T) {
	node := &Node[T]{Vertex: dest}
	node.Next = g.AdjacencyList[src]
	g.AdjacencyList[src] = node

	node = &Node[T]{Vertex: src}
	node.Next = g.AdjacencyList[dest]
	g.AdjacencyList[dest] = node
}

func (g *Graph[T]) RemoveEdge(src, dest T) {
	node := g.AdjacencyList[src]
	if node == nil {
		return
	}

	if node.Vertex == dest {
		g.AdjacencyList[src] = node.Next
		return
	}

	for node.Next != nil {
		if node.Next.Vertex == dest {
			node.Next = node.Next.Next
			return
		}
		node = node.Next
	}
}

func (g *Graph[T]) RemoveVertex(vertex T) {
	for k := range g.AdjacencyList {
		g.RemoveEdge(k, vertex)
	}
	delete(g.AdjacencyList, vertex)
}

func (g *Graph[T]) DFS(start T) []T {
	visited := make(map[T]bool)
	result := []T{}
	g.dfs(start, visited, &result)
	return result
}

func (g *Graph[T]) dfs(vertex T, visited map[T]bool, result *[]T) {
	if visited[vertex] {
		return
	}
	visited[vertex] = true
	*result = append(*result, vertex)

	for node := g.AdjacencyList[vertex]; node != nil; node = node.Next {
		g.dfs(node.Vertex, visited, result)
	}
}
