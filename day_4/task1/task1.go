package task1

import (
	"bufio"
	"log"
	"os"
)

func ConvertToMatrix(r *bufio.Reader) [][]string {
	var matrix [][]string

	for {
		line, _, err := r.ReadLine()

		if len(line) == 0 {
			break
		}

		if err != nil {
			log.Fatal(err)
			break
		}

		var matrixRow []string

		for _, lineChar := range line {
			matrixRow = append(matrixRow, string(lineChar))
		}

		matrix = append(matrix, matrixRow)
	}

	return matrix
}

func SearchedStringSpaceAvailable(matrix *[][]string, rowId int, colId int, searchedText string, direction string) bool {
	n := func() bool {
		return rowId-(len(searchedText)-1) >= 0
	}

	s := func() bool {
		return rowId+(len(searchedText)-1) <= len(*matrix)-1
	}

	e := func() bool {
		return colId+(len(searchedText)-1) <= len((*matrix)[0])-1
	}

	w := func() bool {
		return colId-(len(searchedText)-1) >= 0
	}

	switch direction {
	case "N":
		return n()
	case "S":
		return s()
	case "E":
		return e()
	case "W":
		return w()
	case "NE":
		return n() && e()
	case "NW":
		return n() && w()
	case "SE":
		return s() && e()
	case "SW":
		return s() && w()
	}

	return false
}

func TargetStringIsPresent(matrix *[][]string, rowId int, colId int, searchedText string, direction string) bool {
	if SearchedStringSpaceAvailable(matrix, rowId, colId, searchedText, direction) == false {
		return false
	}

	compareSearchedTextWithinDirection := func(rowIdModifier int, colIdModifier int) bool {
		for j, s := range searchedText {
			modRowId := rowId + (j * rowIdModifier)
			modColId := colId + (j * colIdModifier)

			if (*matrix)[modRowId][modColId] != string(s) {
				return false
			}
		}

		return true
	}

	switch direction {
	case "N":
		return compareSearchedTextWithinDirection(-1, 0)
	case "S":
		return compareSearchedTextWithinDirection(1, 0)
	case "E":
		return compareSearchedTextWithinDirection(0, 1)
	case "W":
		return compareSearchedTextWithinDirection(0, -1)
	case "NE":
		return compareSearchedTextWithinDirection(-1, 1)
	case "NW":
		return compareSearchedTextWithinDirection(-1, -1)
	case "SE":
		return compareSearchedTextWithinDirection(1, 1)
	case "SW":
		return compareSearchedTextWithinDirection(1, -1)
	}

	return true
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

	var total int
	total = 0
	const SEARCHED_TEXT = "XMAS"
	var directions = []string{"N", "S", "E", "W", "NE", "NW", "SE", "SW"}

	matrix := ConvertToMatrix(r)

	for rowId, row := range matrix {
		for colId, _ := range row {
			for _, direction := range directions {
				if TargetStringIsPresent(&matrix, rowId, colId, SEARCHED_TEXT, direction) == true {
					total += 1
				}
			}

			// if rowId == 30 && colId == 94 {
			// 	os.Exit(0)
			// }
			//			os.Exit(0)
		}
	}

	return total
}
