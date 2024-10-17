package main

import (
	"fmt"
	"log"
	"sort"

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
	for name := range rooms {
		if name == "##start" {
			startRoom = name
		} else if name == "##end" {
			endRoom = name
		} else {
			roomNames = append(roomNames, name)
		}
	}

	// Sort the normal rooms alphabetically (you could customize this sorting logic if needed)
	sort.Strings(roomNames)

	fmt.Println("Rooms:")
	// Print start room first (if it exists)
	if startRoom != "" {
		room := rooms[startRoom]
		fmt.Printf("Room %s at coordinates (%d, %d)\n", room.Name, room.X, room.Y)
	}

	// Print all other rooms in alphabetical order
	for _, roomName := range roomNames {
		room := rooms[roomName]
		fmt.Printf("Room %s at coordinates (%d, %d)\n", room.Name, room.X, room.Y)
	}

	// Print end room last (if it exists)
	if endRoom != "" {
		room := rooms[endRoom]
		fmt.Printf("Room %s at coordinates (%d, %d)\n", room.Name, room.X, room.Y)
	}

	fmt.Println("Tunnels:")
	for _, tunnel := range tunnels {
		fmt.Printf("Tunnel between %s and %s\n", tunnel[0], tunnel[1])
	}
}
