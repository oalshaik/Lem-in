package graph

import (
	"fmt"
)

type Graph struct {
	Nodes map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string][]string),
	}
}

func (g *Graph) AddEdge(room1, room2 string) {
	g.Nodes[room1] = append(g.Nodes[room1], room2)
	g.Nodes[room2] = append(g.Nodes[room2], room1)
}

func (g *Graph) DisplayGraph() {
	fmt.Println("Graph representation:")
	for room, neighbors := range g.Nodes {
		fmt.Printf("Room %s is connected to: %v\n", room, neighbors)
	}
}

func (g *Graph) BFS(start, end string) ([]string, bool) {
	queue := [][]string{{start}}

	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {

		path := queue[0]
		queue = queue[1:]

		currentRoom := path[len(path)-1]

		if currentRoom == end {
			return path, true
		}

		for _, neighbor := range g.Nodes[currentRoom] {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}

	}

	return nil, false
}

// FindPaths performs a modified BFS to find multiple shortest paths from start to end
// FindPaths performs a modified BFS to find multiple shortest paths from start to end
// FindPaths performs a modified BFS to find multiple shortest paths from start to end
func (g *Graph) FindPaths(start, end string, maxPaths int) [][]string {
	paths := [][]string{}
	queue := [][]string{{start}}
	roomOccupancy := make(map[int]map[string]bool)

	for len(queue) > 0 && len(paths) < maxPaths {
		path := queue[0]
		queue = queue[1:]

		lastRoom := path[len(path)-1]

		if lastRoom == end {
			if isValidPath(path, roomOccupancy) {
				paths = append(paths, path)
				updateRoomOccupancy(path, roomOccupancy)
			}
			continue
		}

		for _, neighbor := range g.Nodes[lastRoom] {
			if !contains(path, neighbor) {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return paths
}

func isValidPath(path []string, roomOccupancy map[int]map[string]bool) bool {
	for i, room := range path {
		if i == 0 || i == len(path)-1 {
			continue // Skip start and end rooms
		}
		if occupancy, exists := roomOccupancy[i]; exists {
			if occupancy[room] {
				return false // Room is already occupied at this turn
			}
		}
	}
	return true
}

func updateRoomOccupancy(path []string, roomOccupancy map[int]map[string]bool) {
	for i, room := range path {
		if i == 0 || i == len(path)-1 {
			continue // Skip start and end rooms
		}
		if _, exists := roomOccupancy[i]; !exists {
			roomOccupancy[i] = make(map[string]bool)
		}
		roomOccupancy[i][room] = true
	}
}

// Helper function to check if a path contains a room
func contains(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}
