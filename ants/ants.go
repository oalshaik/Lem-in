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

func MoveAntsDynamically(numAnts int, paths [][]string) {
	ants := make([]Ant, numAnts)
	for i := range ants {
		ants[i] = Ant{id: i + 1, pathIndex: i % len(paths), position: 0}
	}

	occupied := make(map[string]bool)
	endRoom := paths[0][len(paths[0])-1]

	for {
		moves := []string{}
		antsAtEnd := 0

		for i := range ants {
			ant := &ants[i]
			currentPath := paths[ant.pathIndex]

			if ant.position == len(currentPath)-1 {
				antsAtEnd++
				continue
			}

			nextRoom := currentPath[ant.position+1]

			if nextRoom == endRoom || !occupied[nextRoom] {
				if nextRoom != endRoom {
					occupied[nextRoom] = true
				}
				moves = append(moves, fmt.Sprintf("L%d-%s", ant.id, nextRoom))
				ant.position++

				if nextRoom == endRoom {
					antsAtEnd++
				}
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}

		if antsAtEnd == numAnts {
			break
		}

		for room := range occupied {
			delete(occupied, room)
		}
	}
}
