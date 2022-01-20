package main

import (
	"fmt"

	"golang-algorithm/graph/standard"
)

func main() {

	graph := standard.NewGraph()

	fmt.Println("------------------start to insert vertex-----------------------")
	graph.InsertVertex(standard.NewVNode("Bill"))
	graph.InsertVertex(standard.NewVNode("Rocket"))
	graph.InsertVertex(standard.NewVNode("Linda"))
	graph.InsertVertex(standard.NewVNode("Shirly"))
	graph.InsertVertex(standard.NewVNode("Bruce"))
	fmt.Println("-------------------finish to insert vertex----------------------")

}
