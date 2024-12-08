package task1

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"

	"github.com/mowshon/iterium"
)

type Frequency struct {
	label     string
	positions []Position
}

type Antinode struct {
	antennaSource *Frequency
}

type Field struct {
	stationedAntenna *Frequency
	antinodes        []Antinode
}

type Position struct {
	rowId, colId int
}

func GetFrequency(label string, frequencies *[]Frequency) *Frequency {
	// Check if already exists
	for i := range *frequencies {
		if (*frequencies)[i].label == label {
			return &(*frequencies)[i]
		}
	}

	newFreq := Frequency{label: label}
	*frequencies = append(*frequencies, newFreq)

	return &(*frequencies)[len(*frequencies)-1]
}

func GetField(stringId string, position Position, frequencies *[]Frequency) Field {
	newField := Field{}

	if stringId == "." {
		return newField
	}

	// Handle frequency entry
	freq := GetFrequency(stringId, frequencies)
	(*freq).positions = append((*freq).positions, position)

	newField.stationedAntenna = freq

	return newField
}

func GetAntennaPairs(freq *Frequency) [][]Position {
	perms, err := iterium.Permutations((*freq).positions, 2).Slice()

	if err != nil {
		log.Fatal("Could not get perms")
		os.Exit(1)
	}

	return perms
}

func ApplyAntennaPairAntinode(area *[][]Field, freq *Frequency, antennaPair []Position) {
	rowDist := antennaPair[0].rowId - antennaPair[1].rowId
	colDist := antennaPair[0].colId - antennaPair[1].colId

	newRowId := antennaPair[0].rowId + rowDist
	newColId := antennaPair[0].colId + colDist

	if newRowId < 0 || newRowId > len(*area)-1 || newColId < 0 || newColId > len((*area)[0])-1 {
		return
	}

	(*area)[newRowId][newColId].antinodes = append((*area)[newRowId][newColId].antinodes, Antinode{antennaSource: freq})
}

func ApplyFrequencyAntinodes(area *[][]Field, freq *Frequency, antennaPairs [][]Position) {
	for _, antennaPair := range antennaPairs {
		ApplyAntennaPairAntinode(area, freq, antennaPair)
	}
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

	frequencies := make([]Frequency, 0)
	area := make([][]Field, 0)

	rowId := 0

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal(err)
			break
		}

		var areaRow []Field

		for colId, rune := range line {
			field := GetField(string(rune), Position{rowId: rowId, colId: colId}, &frequencies)
			areaRow = append(areaRow, field)
		}

		area = append(area, areaRow)

		rowId += 1
	}

	// Loop through all frequency antennas and check all possible antinodes
	for _, freq := range frequencies {
		antennaPairs := GetAntennaPairs(&freq)

		ApplyFrequencyAntinodes(&area, &freq, antennaPairs)
	}

	totalAntinodes := 0

	// Check how many fields have antinodes
	for _, row := range area {
		for _, field := range row {
			if field.antinodes != nil {
				totalAntinodes += 1
			}
		}
	}

	return totalAntinodes
}
