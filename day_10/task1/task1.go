package task1

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

type TrailNode struct {
	value int
	id    int
	n     *TrailNode
	e     *TrailNode
	s     *TrailNode
	w     *TrailNode
}

const MAX_TARGET = 9

func FindAllTrailheads(hikingMap *[][]TrailNode, trailheads *[]*TrailNode) {
	for rowId, row := range *hikingMap {
		for colId, trailNode := range row {
			if trailNode.value == 0 {
				*trailheads = append(*trailheads, &(*hikingMap)[rowId][colId])
			}
		}
	}
}

func DiscoverNeighbours(hikingMap *[][]TrailNode) {
	maxRowId := len(*hikingMap) - 1
	for rowId, row := range *hikingMap {
		maxColId := len(row) - 1
		for colId := range row {
			if rowId-1 >= 0 {
				(*hikingMap)[rowId][colId].n = &(*hikingMap)[rowId-1][colId]
			}

			if colId+1 <= maxColId {
				(*hikingMap)[rowId][colId].e = &(*hikingMap)[rowId][colId+1]
			}

			if rowId+1 <= maxRowId {
				(*hikingMap)[rowId][colId].s = &(*hikingMap)[rowId+1][colId]
			}

			if colId-1 >= 0 {
				(*hikingMap)[rowId][colId].w = &(*hikingMap)[rowId][colId-1]
			}
		}
	}
}

func (t *TrailNode) Neighbours() []*TrailNode {
	var neighbours []*TrailNode

	if t.n != nil {
		neighbours = append(neighbours, t.n)
	}

	if t.e != nil {
		neighbours = append(neighbours, t.e)
	}

	if t.s != nil {
		neighbours = append(neighbours, t.s)
	}

	if t.w != nil {
		neighbours = append(neighbours, t.w)
	}

	return neighbours
}

func (currentTrailNode *TrailNode) ResolveNextTrailNode(currentPath []*TrailNode) [][]*TrailNode {
	var resolvedNewPaths [][]*TrailNode

	// If the current trail is target "height" - mark as resolved fully
	if currentTrailNode.value == MAX_TARGET {
		var resolvedNewPath []*TrailNode
		resolvedNewPath = append(resolvedNewPath, currentPath...)
		resolvedNewPath = append(resolvedNewPath, currentTrailNode)
		resolvedNewPaths = append(resolvedNewPaths, resolvedNewPath)

		return resolvedNewPaths
	}

	neighbours := currentTrailNode.Neighbours()

	var possibleNextTrailNodes []*TrailNode

	for i := range neighbours {
		if neighbours[i].value == currentTrailNode.value+1 {
			possibleNextTrailNodes = append(possibleNextTrailNodes, neighbours[i])
		}
	}

	// If no new possible trail nodes can be found - mark as not resolved
	if len(possibleNextTrailNodes) == 0 {
		return nil
	}

	for i := range possibleNextTrailNodes {
		newTrailPath := append(currentPath, currentTrailNode)

		resolvedNewPaths = append(resolvedNewPaths, possibleNextTrailNodes[i].ResolveNextTrailNode(newTrailPath)...)
	}

	return resolvedNewPaths
}

func GetUniqueTops(foundTrails [][]*TrailNode) []int {
	var uniqueTopIds []int

	for _, foundTrail := range foundTrails {
		topId := foundTrail[len(foundTrail)-1].id

		if !slices.Contains(uniqueTopIds, topId) {
			uniqueTopIds = append(uniqueTopIds, topId)
		}
	}

	return uniqueTopIds
}

func Task1() int {
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

	var hikingMap [][]TrailNode

	trailNodeId := 0

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal(err)
			break
		}

		var hikingMapRow []TrailNode

		for _, rune := range line {
			value, err := strconv.ParseInt(string(rune), 10, 0)

			if err != nil {
				log.Fatal("Could not parse trail node")
				os.Exit(1)
			}

			trailNode := TrailNode{value: int(value), id: trailNodeId}
			hikingMapRow = append(hikingMapRow, trailNode)
			trailNodeId += 1
		}

		hikingMap = append(hikingMap, hikingMapRow)
	}

	var trailheads []*TrailNode

	FindAllTrailheads(&hikingMap, &trailheads)

	DiscoverNeighbours(&hikingMap)

	totalFoundTrails := 0

	for _, trailhead := range trailheads {
		foundTrails := trailhead.ResolveNextTrailNode(nil)

		uniqueTopIds := GetUniqueTops(foundTrails)

		totalFoundTrails += len(uniqueTopIds)
	}

	return totalFoundTrails
}
