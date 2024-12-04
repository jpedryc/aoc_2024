package task2

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

	var total int
	total = 0
	const SEARCHED_TEXT = "XMAS"

	matrix := ConvertToMatrix(r)

	for rowId, row := range matrix {
		for colId, _ := range row {
			// Ignore edges
			if rowId == 0 || rowId == len(matrix)-1 || colId == 0 || colId == len(matrix[0])-1 {
				continue
			}

			isMasOrSam := func(text string) bool {
				return text == "MAS" || text == "SAM"
			}

			if matrix[rowId][colId] == "A" {
				if isMasOrSam(matrix[rowId-1][colId-1]+matrix[rowId][colId]+matrix[rowId+1][colId+1]) && isMasOrSam(matrix[rowId+1][colId-1]+matrix[rowId][colId]+matrix[rowId-1][colId+1]) {
					total += 1
				}
			}
		}
	}

	return total
}
