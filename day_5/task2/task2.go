package task2

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"slices"
	"sort"
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

func RuleForPagesValid(minPage *int, maxPages *[]int, pages []int) bool {
	minIndex := slices.Index(pages, *minPage)

	if minIndex == -1 {
		return true
	}

	var maxPageIndexes []int
	for _, maxPage := range *maxPages {
		maxPageIndexes = append(maxPageIndexes, slices.Index(pages, maxPage))
	}

	maxPageIndexesClean := slices.DeleteFunc(maxPageIndexes, func(idx int) bool { return idx == -1 })

	if len(maxPageIndexesClean) == 0 {
		return true
	}

	sort.Ints(maxPageIndexesClean)

	return minIndex < maxPageIndexesClean[0]
}

func RulesForPagesValid(rules map[int][]int, pages []int) bool {
	for minPage, maxPages := range rules {
		if RuleForPagesValid(&minPage, &maxPages, pages) == false {
			return false
		}
	}

	return true
}

func ApplyRuleToPages(minPage *int, maxPages *[]int, pages []int) []int {
	minIndex := slices.Index(pages, *minPage)

	if minIndex == -1 {
		return pages
	}

	var maxPageIndexes []int
	for _, maxPage := range *maxPages {
		maxPageIndexes = append(maxPageIndexes, slices.Index(pages, maxPage))
	}

	maxPageIndexesClean := slices.DeleteFunc(maxPageIndexes, func(idx int) bool { return idx == -1 })

	if len(maxPageIndexesClean) == 0 {
		return pages
	}

	sort.Ints(maxPageIndexesClean)

	if minIndex < maxPageIndexesClean[0] {
		return pages
	}

	modifiedPages := slices.Clone(pages)[:maxPageIndexesClean[0]]
	modifiedPages = append(modifiedPages, pages[minIndex])

	rest := slices.Clone(pages)[maxPageIndexesClean[0]:]
	third := slices.DeleteFunc(append([]int{}, rest...), func(v int) bool { return v == pages[minIndex] })

	modifiedPages = append(modifiedPages, third...)

	return modifiedPages
}

func ApplyRulesToPages(rules map[int][]int, pages []int) []int {
	var reorderedPages []int
	reorderedPages = slices.Clone(pages)

	for minPage, maxPages := range rules {
		reorderedPages = ApplyRuleToPages(&minPage, &maxPages, reorderedPages)
	}

	if slices.Equal(reorderedPages, pages) {
		return []int{}
	}

	return reorderedPages
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

	var rules map[int][]int
	rules = map[int][]int{}
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
			ruleEntries := ConvertToInt(ruleEntriesSplitted)
			rules[ruleEntries[0]] = append(rules[ruleEntries[0]], ruleEntries[1])
		} else if strings.Contains(string(line), ",") {
			pageEntriesSplitted := strings.Split(string(line), ",")
			pages = append(pages, ConvertToInt(pageEntriesSplitted))
		}
	}

	var middlePageSum int
	middlePageSum = 0

	for _, pageEntry := range pages {
		if RulesForPagesValid(rules, pageEntry) == false {
			reorderedPageEntry := ApplyRulesToPages(rules, pageEntry)

			middlePageSum += reorderedPageEntry[len(reorderedPageEntry)/2]
		}
	}

	return middlePageSum
}
