package main

import (
	"day6/task1"
	"day6/task2"
	"fmt"
)

func main() {
	fmt.Println("Day 6 Results")
	task1Result := task1.Task1()
	fmt.Println("Task1 result:", task1Result) // 4656
	task2Result := task2.Task2()
	fmt.Println("Task2 result:", task2Result) // 1575
}
