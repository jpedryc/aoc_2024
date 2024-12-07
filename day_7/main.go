package main

import (
	"day7/task1"
	"day7/task2"
	"fmt"
)

func main() {
	fmt.Println("Day 7 Results")
	task1Result := task1.Task1()
	fmt.Println("Task1 result:", task1Result) // 20665830408335
	task2Result := task2.Task2()
	fmt.Println("Task2 result:", task2Result) // 354060705047464
}
