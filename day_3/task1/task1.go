package task1

import (
	"bufio"
	"fmt"
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

	//regex := regexp.MustCompile(`mul\(\d{1-3},\d{1-3}\)`)
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	r := bufio.NewReader(file)

	var total int
	total = 0

	for {
		line, _, err := r.ReadLine()

		if len(line) == 0 {
			break
		}

		if err != nil {
			log.Fatal(err)
			break
		}

		mulStrings := regex.FindAllStringSubmatch(string(line[:]), -1)

		for _, submatch := range mulStrings {
			first, err := strconv.ParseInt(submatch[1], 10, 0)

			if err != nil {
				log.Fatal(err)
				break
			}
			second, err := strconv.ParseInt(submatch[2], 10, 0)

			if err != nil {
				log.Fatal(err)
				break
			}

			total += int(first) * int(second)
		}

		fmt.Println(mulStrings)
	}

	return total
}
