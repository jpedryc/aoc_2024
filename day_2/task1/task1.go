package task1

import (
	"bufio"
	"log"
	"os"
	"regexp"
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

	regex := regexp.MustCompile(`\s+`)

	r := bufio.NewReader(file)

	var previousVal int64
	var valInt int64
	var increasing bool
	safeReports := 0
	var safe bool

	for {
		line, _, err := r.ReadLine()

		safe = true
		if len(line) > 0 {
			results := regex.Split(string(line[:]), -1)

			for idx, val := range results {
				valInt, err = strconv.ParseInt(val, 10, 0)

				if err != nil {
					log.Fatal(err)
				}

				if idx == 0 {
					previousVal = valInt
					continue
				}

				if idx == 1 {
					if valInt > previousVal {
						increasing = true
					} else {
						increasing = false
					}
				}

				diff := valInt - previousVal

				if increasing == true {
					if diff >= 1 && diff <= 3 {
						previousVal = valInt
						continue
					} else {
						safe = false
						break
					}
				} else {
					if diff <= -1 && diff >= -3 {
						previousVal = valInt
						continue
					} else {
						safe = false
						break
					}
				}
			}

			if safe == true {
				safeReports += 1
			} else {
			}
		}

		if err != nil {
			break
		}
	}

	return safeReports
}
