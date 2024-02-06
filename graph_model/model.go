package graph_model

import "fmt"

/*
Mi idea de grafo es que sea un grafo ponderado
la primer key del map representa el nodo, la segunda key del siguiente map representa el nodo con el cual comparte
una arista y el int final representa el peso
*/
type Graph struct {
	Nodes map[int]map[int]int
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]map[int]int),
	}
}

func (it *Graph) AddEdge(nodeOrigin, nodeDestiny, weight int) {
	//Si el nodo de origen no existe lo crea
	if _, ok := it.Nodes[nodeOrigin]; !ok {
		it.Nodes[nodeOrigin] = make(map[int]int)
	}

	//Si el nodo de destino no existe lo crea
	if _, ok := it.Nodes[nodeDestiny]; !ok {
		it.Nodes[nodeDestiny] = make(map[int]int)
	}

	//agrega la arista con su peso
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
	isFirst := true
	for node, edges := range it.Nodes {
		fmt.Printf("Node %d -> [", node)
		for nodeEdge, weight := range edges {
			if !isFirst {
				fmt.Printf(", ")
			}
			fmt.Printf(" ( edge [%d, %d], weight: %d) ", node, nodeEdge, weight)
			isFirst = false
		}
		fmt.Print("]")
		isFirst = true
		fmt.Println()
	}
}

func (it *Graph) DeleteNode(nodeToDelete int) bool {
	if _, ok := it.Nodes[nodeToDelete]; ok {
		delete(it.Nodes, nodeToDelete)

		//Al no existir el nodo es importante eliminar de los nodos restantes las aristas que comarten
		//Ya que estas tampoco existirían
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
