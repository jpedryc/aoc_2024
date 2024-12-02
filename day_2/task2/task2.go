package task2

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func ConvertToInt(inputData []string) []int64 {
	var outputData []int64

	for _, s := range inputData {
		i, err := strconv.ParseInt(s, 10, 0)

		if err != nil {
			log.Fatal(err)
		}

		outputData = append(outputData, i)
	}

	return outputData
}

func IsIncreasing(inputData []int64) bool {
	var incCnt = 0
	var decCnt = 0
	var previous int64

	for i, v := range inputData {
		if i == 0 {
			previous = v
			continue
		}

		if v > previous {
			incCnt += 1
		} else if v < previous {
			decCnt += 1
		}

		previous = v
	}

	return incCnt > decCnt
}

func LocationListValid(locationList []int64, increasing bool) bool {
	var previousValue *int64
	previousValue = nil

	for _, value := range locationList {
		if previousValue == nil {
			previousValue = &value
			continue
		}

		diff := value - *previousValue

		if increasing == false {
			diff *= -1
		}

		if diff < 1 || diff > 3 {
			return false
		}

		previousValue = &value
	}

	return true
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

	regex := regexp.MustCompile(`\s+`)

	r := bufio.NewReader(file)

	var safeReports int
	safeReports = 0

	var safe bool

	for {
		line, _, err := r.ReadLine()

		if len(line) == 0 {
			break
		}

		if err != nil {
			log.Fatal(err)
			break
		}

		locsStrings := regex.Split(string(line[:]), -1)

		locs := ConvertToInt(locsStrings[:])

		increasing := IsIncreasing(locs[:])
		if increasing {
		} else {
		}

		safe = true

		// Check original list
		safe = LocationListValid(locs, increasing)

		if safe == false {
			// Check modified lists
			for j := 0; j < len(locs); j++ {
				clonedList := slices.Clone(locs)
				modifiedList := slices.Delete(clonedList, j, j+1)
				safe = LocationListValid(modifiedList, increasing)

				if safe {
					break
				}
			}
		}

		if safe == true {
			safeReports += 1
		} else {
		}
	}

	return safeReports
}
