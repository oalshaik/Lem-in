package ants

import (
	"fmt"
	"strings"
)

// Ant represents each ant's state
type Ant struct {
	id        int
	pathIndex int
	position  int
}

// MoveAntsDynamically simulates the movement of ants dynamically considering room occupancy
func MoveAntsDynamically(numAnts int, paths [][]string) {
	ants := make([]Ant, numAnts)
	// Assign ants to paths
	for i := range ants {
		ants[i] = Ant{id: i + 1, pathIndex: i % len(paths), position: 0}
	}

	occupied := make(map[string]bool) // Track occupied rooms, except for start and end rooms

	for {
		moves := []string{}
		allArrived := true

		for i := range ants {
			ant := &ants[i]
			currentPath := paths[ant.pathIndex]

			// If ant has reached the end, skip it
			if ant.position == len(currentPath)-1 {
				continue
			}

			allArrived = false
			nextRoom := currentPath[ant.position+1]

			// Allow ant to move if the next room is not occupied (except for start and end rooms)
			if nextRoom != currentPath[0] && !occupied[nextRoom] {
				occupied[nextRoom] = true
				moves = append(moves, fmt.Sprintf("L%d-%s", ant.id, nextRoom))
				ant.position++ // Move the ant to the next room
			}
		}

		// Print moves for this turn
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}

		// Stop if all ants have arrived at their destination
		if allArrived {
			break
		}

		// Clear the occupied map after each turn (except for start and end rooms)
		for room := range occupied {
			delete(occupied, room)
		}
	}
}
