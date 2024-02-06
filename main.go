package main

import (
	"fmt"
	"graphs/graph_model"
)

func main() {

	//graphExample1()
	//graphExample2()
	//graphExample3()
	//graphExample4()
}

func graphExample1() {
	graph := graph_model.NewGraph()
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(2, 3, 5)
	graph.AddEdge(2, 4, 5)
	graph.AddEdge(2, 5, 3)
	graph.AddEdge(3, 4, 1)
	graph.AddEdge(4, 5, 7)
	graph.AddEdge(4, 6, 9)

	graph.PrintGraph()
}

// path
func graphExample2() {
	graph := graph_model.NewGraph()
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(2, 3, 5)
	graph.AddEdge(3, 4, 5)
	graph.AddEdge(4, 5, 3)
	graph.AddEdge(5, 6, 1)
	graph.AddEdge(6, 7, 7)

	graph.PrintGraph()
}

// tree
func graphExample3() {
	graph := graph_model.NewGraph()
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 4, 5)
	graph.AddEdge(2, 5, 3)
	graph.AddEdge(3, 6, 1)
	graph.AddEdge(3, 7, 7)

	graph.PrintGraph()
}

func graphExample4() {
	graph := graph_model.NewGraph()
	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(33)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 4, 5)
	graph.AddEdge(2, 5, 3)
	graph.AddEdge(3, 6, 1)
	graph.AddEdge(3, 7, 7)
	graph.PrintGraph()

	fmt.Println("delete node 33")
	graph.DeleteNode(33)
	graph.PrintGraph()

	fmt.Println("add edge in node 1 to 2, with weight 2")
	graph.AddEdge(1, 2, 2)
	graph.PrintGraph()

	fmt.Println("add edge in node 3 to 2, with weight 10")
	graph.AddEdge(3, 2, 10)
	graph.PrintGraph()

	fmt.Println("delete node 2")
	graph.DeleteNode(2)
	graph.PrintGraph()
}
