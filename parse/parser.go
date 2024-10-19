package parse

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X    int
	Y    int
}

func ParseFile(filename string) (int, map[string]Room, [][2]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numAnts int
	rooms := make(map[string]Room)
	var tunnels [][2]string

	var startRoom, endRoom string
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if strings.HasPrefix(line, "#") {
			if line == "##start" || line == "##end" {
				if line == "##start" {
					startRoom = "start"
				} else {
					endRoom = "end"
				}
			}
			continue
		}
		if lineNum == 1 {
			numAnts, err = strconv.Atoi(line)
			if err != nil || numAnts <= 0 {
				return 0, nil, nil, fmt.Errorf("invalid number of ants")
			}
			continue
		}
		if isRoom(line) {
			parts := strings.Split(line, " ")
			if len(parts) != 3 {
				return 0, nil, nil, fmt.Errorf("invalid room format at line %d", lineNum)
			}
			name := parts[0]
			x, err := strconv.Atoi(parts[1])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("invalid x coordinate at line %d", lineNum)
			}
			y, err := strconv.Atoi(parts[2])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("invalid y coordinate at line %d", lineNum)
			}

			rooms[name] = Room{Name: name, X: x, Y: y}

			if startRoom == "start" {
				startRoom = name
				rooms["##start"] = Room{Name: name, X: x, Y: y}
			} else if endRoom == "end" {
				endRoom = name
				rooms["##end"] = Room{Name: name, X: x, Y: y}
			}

			continue
		}

		if isTunnel(line) {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return 0, nil, nil, fmt.Errorf("invalid tunnel format at line %d", lineNum)
			}
			tunnels = append(tunnels, [2]string{parts[0], parts[1]})
		} else {
			return 0, nil, nil, fmt.Errorf("unknown format at line %d", lineNum)
		}

	}

	if _, ok := rooms["##start"]; !ok {
		return 0, nil, nil, fmt.Errorf("missing start room")
	}
	if _, ok := rooms["##end"]; !ok {
		return 0, nil, nil, fmt.Errorf("missing end room")
	}

	return numAnts, rooms, tunnels, nil

}

func isRoom(line string) bool {
	parts := strings.Split(line, " ")
	return len(parts) == 3 && !strings.HasPrefix(parts[0], "L") && !strings.HasPrefix(parts[0], "#")
}

func isTunnel(line string) bool {
	return strings.Count(line, "-") == 1
}
