package task2

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

	var maxId int64
	maxId = GetElementInt(disk.Back())

	// Move file block from the back to front
	for endElement := disk.Back(); endElement != nil; endElement = endElement.Prev() {

		endInt := GetElementInt(endElement)

		// We want to move only positive values and make sure we're going according to IDs (desc)
		if endInt == int64(-1) || endInt != maxId {
			continue
		}

		// Save the initial found element
		endPosStarter := endElement

		// Start counting the chain
		currentFileIdChain := 1

		// Check how far the chaing goes
		for {
			endElement = endElement.Prev()

			if endElement == nil {
				break
			}

			endInt := GetElementInt(endElement)

			if endInt == maxId {
				currentFileIdChain += 1
			} else {
				break
			}
		}

		maxId -= 1

		if endElement == nil {
			break
		}

		// Revert to keep loop INC statement correct
		endElement = endElement.Next()

		for startElement := disk.Front(); startElement != endElement; startElement = startElement.Next() {
			startInt := GetElementInt(startElement)

			// We have to find next possible empty space
			if startInt != int64(-1) {
				continue
			}

			// Save the initial found space
			startPosStarter := startElement

			// Start counting the chain
			currentSpaceIdChain := 1

			// Check how far the chain goes
			for {
				startElement = startElement.Next()

				if startElement == nil {
					break
				}

				startInt := GetElementInt(startElement)

				if startInt == -1 {
					currentSpaceIdChain += 1
				} else {
					break
				}
			}

			// Revert to keep loop INC statement correct
			startElement = startElement.Prev()

			if startElement == nil {
				break
			}

			if currentSpaceIdChain >= currentFileIdChain {
				// Replace found spaces with file IDs
				for j := 0; j < currentFileIdChain; j++ {
					tmpStartValue := startPosStarter.Value
					startPosStarter.Value = endPosStarter.Value
					endPosStarter.Value = tmpStartValue

					startPosStarter = startPosStarter.Next()
					endPosStarter = endPosStarter.Prev()

					if startPosStarter == nil || endPosStarter == nil {
						break
					}
				}

				break
			}
		}
	}

	var total uint64
	total = 0

	var pos int64
	pos = 0

	for head := disk.Front(); head != nil; head = head.Next() {
		eInt := GetElementInt(head)

		if eInt > int64(-1) {
			total += uint64(pos * eInt)
		}

		pos += 1
	}

	return total
}
