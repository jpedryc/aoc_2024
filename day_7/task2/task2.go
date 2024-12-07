package task2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/mowshon/iterium"
)

type Calibration struct {
	result    uint64
	arguments []int
}

func getCalibrationFromLine(line string, calRegex *regexp.Regexp, argsRegex *regexp.Regexp) Calibration {
	calibrationElements := calRegex.FindAllStringSubmatch(string(line), -1)

	result := calibrationElements[0][1]

	res, err := strconv.ParseUint(result, 10, 64)

	if err != nil {
		log.Fatalln("Error while parsing result from file")
		os.Exit(1)
	}

	argsSplitted := argsRegex.Split(calibrationElements[0][2], -1)

	var args []int
	for _, arg := range argsSplitted {
		a, err := strconv.ParseUint(arg, 10, 0)

		if err != nil {
			log.Fatalln("Error while parsing args from file")
			os.Exit(1)
		}

		args = append(args, int(a))
	}

	return Calibration{
		result:    res,
		arguments: args,
	}
}

func GetOpsCombinations(ops *[]string, length int) [][]string {
	combis, err := iterium.Product(*ops, length).Slice()

	if err != nil {
		log.Fatalln("Error while getting combis")
		os.Exit(1)
	}

	return combis
}

func ValidCalibrationOperation(calibration *Calibration, operationCombis *[]string) bool {
	var calcResult uint64
	calcResult = 0

	for argIdx, arg := range calibration.arguments {
		if argIdx == 0 {
			calcResult += uint64(arg)
			continue
		}

		switch (*operationCombis)[argIdx-1] {
		case "+":
			calcResult += uint64(arg)
		case "*":
			calcResult *= uint64(arg)
		case "|":
			calcResultStr := fmt.Sprint(calcResult)
			argStr := fmt.Sprint(arg)
			newTotalStr := fmt.Sprintf("%s%s", calcResultStr, argStr)
			newCalcResult, err := strconv.ParseUint(newTotalStr, 10, 64)

			if err != nil {
				log.Fatalln("Error while combining strings")
				os.Exit(1)
			}

			calcResult = newCalcResult
		}

		// Because the operation are additive, we can jump out as soon as the calculated result is higher than the target
		if calcResult > calibration.result {
			return false
		}
	}

	return calcResult == calibration.result
}

func ValidCalibrationOperations(calibration *Calibration, operationCombis [][]string) int {
	validCalibrationOperationCntr := 0

	for _, opsCombi := range operationCombis {
		if ValidCalibrationOperation(calibration, &opsCombi) == true {
			validCalibrationOperationCntr += 1
		}
	}

	return validCalibrationOperationCntr
}

func Task2() uint64 {
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

	calRegex := regexp.MustCompile(`(\d+): (.*)`)
	argsRegex := regexp.MustCompile(`\s`)

	var calibrations []Calibration

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal(err)
			break
		}

		calibrations = append(calibrations, getCalibrationFromLine(string(line), calRegex, argsRegex))
	}

	var validCalibrationsCntr uint64
	validCalibrationsCntr = 0
	operations := []string{"+", "*", "|"}

	for _, cal := range calibrations {
		opsCombis := GetOpsCombinations(&operations, len(cal.arguments)-1)

		validOps := ValidCalibrationOperations(&cal, opsCombis)

		if validOps > 0 {
			validCalibrationsCntr += cal.result
		}
	}

	return validCalibrationsCntr
}
