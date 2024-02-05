package graph_model

import "fmt"

type Graph struct {
	Nodes map[int]map[int]int
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]map[int]int),
	}
}

func (it *Graph) AddEdge(nodeOrigin, nodeDestiny, weight int) {
	if _, ok := it.Nodes[nodeOrigin]; !ok {
		it.Nodes[nodeOrigin] = make(map[int]int)
	}
	it.Nodes[nodeOrigin][nodeDestiny] = weight
}

func (it *Graph) AddNode(node int) bool {
	if _, ok := it.Nodes[node]; !ok {
		it.Nodes[node] = make(map[int]int)
		return !ok
	}

	return false
}

func (it *Graph) PrintGraph() {
	for node, edges := range it.Nodes {
		fmt.Printf("Node %d -> ", node)
		for nodeEdge, weight := range edges {
			fmt.Printf("( edge [%d, %d], weight: %d) ", node, nodeEdge, weight)
		}
		fmt.Println()
	}
}

func (it *Graph) DeleteNode(nodeToDelete int) bool {
	if _, ok := it.Nodes[nodeToDelete]; ok {
		delete(it.Nodes, nodeToDelete)

		for node, edges := range it.Nodes {
			for nodeEdge, _ := range edges {
				if nodeEdge == nodeToDelete {
					delete(it.Nodes[node], nodeToDelete)
				}
			}
		}

		return ok
	}

	return false
}
