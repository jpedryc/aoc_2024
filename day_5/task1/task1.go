package task1

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ConvertToInt(inputData []string) []int {
	var outputData []int

	for _, s := range inputData {
		i, err := strconv.ParseInt(s, 10, 0)

		if err != nil {
			log.Fatal(err)
		}

		outputData = append(outputData, int(i))
	}

	return outputData
}

func RuleForPagesValid(rule []int, pages []int) bool {
	minIndex := slices.Index(pages, rule[0])

	if minIndex == -1 {
		return true
	}

	maxIndex := slices.Index(pages, rule[1])

	if maxIndex == -1 {
		return true
	}

	return minIndex < maxIndex
}

func RulesForPagesValid(rules [][]int, pages []int) bool {
	for _, rule := range rules {
		if RuleForPagesValid(rule, pages) == false {
			return false
		}
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

	var rules [][]int
	var pages [][]int

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal(err)
			break
		}

		if strings.Contains(string(line), "|") {
			ruleEntriesSplitted := strings.Split(string(line), "|")
			rules = append(rules, ConvertToInt(ruleEntriesSplitted))
		} else if strings.Contains(string(line), ",") {
			pageEntriesSplitted := strings.Split(string(line), ",")
			pages = append(pages, ConvertToInt(pageEntriesSplitted))
		}
	}

	var middlePageSum int
	middlePageSum = 0

	for _, pageEntry := range pages {
		if RulesForPagesValid(rules, pageEntry) == true {
			middlePageSum += pageEntry[len(pageEntry)/2]
		}
	}

	return middlePageSum
}
