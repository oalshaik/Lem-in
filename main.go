package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/oalshaik/Lem-in/ants"
	"github.com/oalshaik/Lem-in/graph"
	"github.com/oalshaik/Lem-in/parse"
)

func main() {
	filename := "testdata//test0.text"

	numAnts, rooms, tunnels, err := parse.ParseFile(filename)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}

	fmt.Printf("Number of Ants: %d\n", numAnts)

	// Create a slice to hold room names
	var roomNames []string

	// Separate start, end, and normal rooms
	var startRoom, endRoom string
	for name, room := range rooms {
		if name == "##start" {
			startRoom = room.Name // Map to the actual start room name
		} else if name == "##end" {
			endRoom = room.Name // Map to the actual end room name
		} else {
			roomNames = append(roomNames, name)
		}
	}

	// Sort the normal rooms alphabetically
	sort.Strings(roomNames)

	fmt.Println("Rooms:")
	if startRoom != "" {
		room := rooms[startRoom]
		fmt.Printf("Room %s at coordinates (%d, %d)\n", room.Name, room.X, room.Y)
	}
	for _, roomName := range roomNames {
		room := rooms[roomName]
		fmt.Printf("Room %s at coordinates (%d, %d)\n", room.Name, room.X, room.Y)
	}
	if endRoom != "" {
		room := rooms[endRoom]
		fmt.Printf("Room %s at coordinates (%d, %d)\n", room.Name, room.X, room.Y)
	}

	// Create the graph representation
	colonyGraph := graph.NewGraph()

	// Add tunnels (edges) to the graph
	for _, tunnel := range tunnels {
		room1, room2 := tunnel[0], tunnel[1]
		colonyGraph.AddEdge(room1, room2)
	}

	fmt.Println("Tunnels (Graph):")
	colonyGraph.DisplayGraph()

	// Find multiple paths from start to end
	maxPaths := 10 // Adjust if needed
	paths := colonyGraph.FindPaths(startRoom, endRoom, maxPaths)

	if len(paths) == 0 {
		fmt.Println("No path found from start to end.")
	} else {
		fmt.Println("Found paths:", paths)
		ants.MoveAntsDynamically(numAnts, paths) // Use the dynamic ant movement
	}
}
