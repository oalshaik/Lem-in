package graph

import (
	"fmt"
)

type Graph struct{
	Nodes map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(make[string][]string)
	}
}

func (g *Graph) AddEdge(room1, room2 string) {
	g.Nodes[room1] = append(g.Nodes[room1], room2)
	g.Nodes[room2] = append(g.Nodes[room2], room1)
}

func (g *Graph) DsiplayGraph(){
	fmt.Println("Graph representation:")
	for room, neighbors := range g.Nodes {
		fmt.Printf("Room %s is connected to: %v\n", room, neighbors)
	}
}