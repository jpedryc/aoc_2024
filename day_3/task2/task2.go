package task2

import (
	"bufio"
	"fmt"
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

	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(do\(\))|(don\'t\(\))`)

	r := bufio.NewReader(file)

	var total int
	total = 0

	var mulEnabled bool
	mulEnabled = true

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

			switch submatchType := submatch[0][0:3]; submatchType {
			case "mul":
				fmt.Println("Handling mul(...)")
				if mulEnabled == false {
					continue
				}
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
			case "do(":
				fmt.Println("Handling do()")
				mulEnabled = true
			case "don":
				fmt.Println("Handling don't()")
				mulEnabled = false
			}
		}
	}

	return total
}
