package main

import (
	"day1/task1"
	"day1/task2"
	"fmt"
)

func main() {
	fmt.Println("Day 1 Results")
	task1Result := task1.Task1()
	fmt.Println("Task1 result:", task1Result) // 1341714
	task2Result := task2.Task2()
	fmt.Println("Task2 result:", task2Result) // 27384707
}
