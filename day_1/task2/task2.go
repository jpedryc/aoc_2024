package task2

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

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

	freq := make(map[int64]int)

	for _, leftVal := range leftLocationIds {
		for _, rightVal := range rightLocationIds {
			if leftVal == rightVal {
				freq[leftVal] += 1
			}
		}
	}

	sum := 0

	for leftVal, freqCnt := range freq {
		sum += int(leftVal) * freqCnt
	}

	return sum
}
