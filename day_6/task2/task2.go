package task2

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

type Field struct {
	blocked bool
	visited bool
}

type GuardPosition struct {
	rowId     int
	colId     int
	direction string
}

func RotateGuard(guardPosition *GuardPosition) {
	switch (*guardPosition).direction {
	case "N":
		(*guardPosition).direction = "E"
	case "E":
		(*guardPosition).direction = "S"
	case "S":
		(*guardPosition).direction = "W"
	case "W":
		(*guardPosition).direction = "N"
	}
}

func MoveGuard(area *[][]Field, guardPosition *GuardPosition, rowModifier int, colModifier int) (bool, error) {
	maxRowId := len(*area) - 1
	maxColId := len((*area)[0]) - 1

	modRowId := guardPosition.rowId + rowModifier
	modColId := guardPosition.colId + colModifier

	// Check if not out of bounds
	if modRowId < 0 || modRowId > maxRowId || modColId < 0 || modColId > maxColId {
		return false, errors.New("Guard left the area")
	}

	nextField := &((*area)[modRowId][modColId])

	// Rotate if blocked
	if (*nextField).blocked == true {
		RotateGuard(guardPosition)
		return false, nil
	}

	visited := false

	// Check if visited
	if nextField.visited != true {
		nextField.visited = true
		visited = true
	}

	// Move the guard
	guardPosition.rowId = modRowId
	guardPosition.colId = modColId

	return visited, nil
}

func GuardVisited(guardPosition *GuardPosition, visitedGuardPositions *[]GuardPosition) bool {
	for _, visitedGuardPosition := range *visitedGuardPositions {
		if visitedGuardPosition.rowId == (*guardPosition).rowId && visitedGuardPosition.colId == (*guardPosition).colId && visitedGuardPosition.direction == (*guardPosition).direction {
			return true
		}
	}

	return false
}

func GuardLooped(area *[][]Field, guardPosition *GuardPosition) bool {
	// We will persist the visited locations and current direction of the guard - if the guard was already on a field
	// and in the same direction - he's in a loop
	var visitedGuardPositions []GuardPosition

	var visited bool
	var err error

	for {
		// Check next possible direction
		switch guardPosition.direction {
		case "N":
			visited, err = MoveGuard(area, guardPosition, -1, 0)
		case "E":
			visited, err = MoveGuard(area, guardPosition, 0, 1)
		case "S":
			visited, err = MoveGuard(area, guardPosition, 1, 0)
		case "W":
			visited, err = MoveGuard(area, guardPosition, 0, -1)
		}

		if err != nil {
			// Guard left the area
			return false
		}

		if GuardVisited(guardPosition, &visitedGuardPositions) == true {
			return true
		}

		if visited == true {
			visitedGuardPositions = append(visitedGuardPositions, GuardPosition{rowId: (*guardPosition).rowId, colId: (*guardPosition).colId, direction: (*guardPosition).direction})
		}
	}
}

func Task2() int {
	file, err := os.Open("input1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(file)

	var area [][]Field
	area = make([][]Field, 130)

	var guardPosition GuardPosition

	lineReadCntr := 0

	// Create area of fields
	for {
		line, _, err := r.ReadLine()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal(err)
			break
		}

		for lineCharIdx, char := range line {
			if char == '.' {
				area[lineReadCntr] = append(area[lineReadCntr], Field{blocked: false, visited: false})
			} else if char == '#' {
				area[lineReadCntr] = append(area[lineReadCntr], Field{blocked: true, visited: false})
			} else if char == '^' {
				area[lineReadCntr] = append(area[lineReadCntr], Field{blocked: false, visited: true})
				guardPosition = GuardPosition{rowId: lineReadCntr, colId: lineCharIdx, direction: "N"}
			} else {
				os.Exit(1)
			}
		}

		lineReadCntr += 1
	}

	guardLoopedCntr := 0

	// Check an area with a new obstraction
	for rowId, row := range area {
		for colId, field := range row {
			initialGuardPosition := GuardPosition{rowId: guardPosition.rowId, colId: guardPosition.colId, direction: guardPosition.direction}

			// Ignore fields where there are already blockers or the guards initial position
			if field.blocked == true || (rowId == initialGuardPosition.rowId && colId == initialGuardPosition.colId) {
				continue
			}

			newArea := make([][]Field, len(area))
			for naIdx := range area {
				newArea[naIdx] = make([]Field, len(area[0]))
				copy(newArea[naIdx], area[naIdx])
			}

			newArea[rowId][colId].blocked = true

			looped := GuardLooped(&newArea, &initialGuardPosition)

			if looped == true {
				guardLoopedCntr += 1
			} else {
			}
		}
	}

	return guardLoopedCntr
}
