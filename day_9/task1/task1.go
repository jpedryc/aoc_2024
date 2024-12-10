package task1

import (
	"bufio"
	"container/list"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

type FileBlock struct {
	fileId int64
	length int
}

type SpaceBlock struct {
	length int
}

func GetElementInt(e *list.Element) int64 {
	v, ok := e.Value.(int64)

	if ok == false {
		log.Fatal("Something went wrong while getting value")
		os.Exit(1)
	}

	return v
}

func Task1() uint64 {
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

	diskMap := list.New()

	isFile := true
	var fileId int64
	fileId = 0

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal(err)
			break
		}

		for _, rune := range line {
			v, err := strconv.ParseInt(string(rune), 10, 0)

			if err != nil {
				log.Fatal("Could not parse to int")
				os.Exit(1)
			}

			// Create Block structs of found disk map
			if isFile {
				diskMap.PushBack(FileBlock{fileId: fileId, length: int(v)})
				fileId += 1
			} else {
				diskMap.PushBack(SpaceBlock{length: int(v)})
			}

			isFile = !isFile
		}
	}

	// Create the real disk image
	disk := list.New()

	for e := diskMap.Front(); e != nil; e = e.Next() {
		switch v := e.Value.(type) {
		case FileBlock:
			fileBlock := v

			for i := 0; i < fileBlock.length; i++ {
				disk.PushBack(fileBlock.fileId)
			}
		case SpaceBlock:
			spaceBlock := v
			for i := 0; i < spaceBlock.length; i++ {
				disk.PushBack(int64(-1))
			}
		default:
			log.Fatal("Something weird happen")
			os.Exit(1)
		}
	}

	// Move file block from the back to front
	for endElement := disk.Back(); endElement != nil; endElement = endElement.Prev() {

		endInt := GetElementInt(endElement)

		// We want to move only positive values
		if endInt == int64(-1) {
			continue
		}

		for startElement := disk.Front(); startElement != endElement; startElement = startElement.Next() {
			startInt := GetElementInt(startElement)

			// We have to find next possible empty space
			if startInt != int64(-1) {
				continue
			}

			tmpStartValue := startElement.Value
			startElement.Value = endElement.Value
			endElement.Value = tmpStartValue

			break
		}
	}

	var total uint64
	total = 0

	var pos int64
	pos = 0

	for head := disk.Front(); head != nil; head = head.Next() {
		eInt := GetElementInt(head)

		if eInt > -1 {
			total += uint64(pos * eInt)
		}

		pos += 1
	}

	return total
}
