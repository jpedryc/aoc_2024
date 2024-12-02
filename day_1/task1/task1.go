package task1

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

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

	var leftLocationIds []int64
	var rightLocationIds []int64

	regex := regexp.MustCompile(`\s+`)

	r := bufio.NewReader(file)

	var rightConv int64
	var leftConv int64

	for {
		line, _, err := r.ReadLine()

		if len(line) > 0 {
			result := regex.Split(string(line[:]), -1)

			leftConv, err = strconv.ParseInt(result[0], 10, 0)

			if err != nil {
				log.Fatal(err)
			}

			rightConv, err = strconv.ParseInt(result[1], 10, 0)

			if err != nil {
				log.Fatal(err)
			}

			leftLocationIds = append(leftLocationIds, leftConv)
			rightLocationIds = append(rightLocationIds, rightConv)
		}

		if err != nil {
			break
		}
	}

	slices.Sort(leftLocationIds)
	slices.Sort(rightLocationIds)

	total := 0

	for idx, leftVal := range leftLocationIds {
		distRaw := leftVal - rightLocationIds[idx]

		if distRaw < 0 {
			total += int(distRaw) * -1
		} else {
			total += int(distRaw)
		}
	}

	return total
}
