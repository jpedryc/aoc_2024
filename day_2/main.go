package main

import (
	"day2/task1"
	"day2/task2"
	"fmt"
)

func main() {
	fmt.Println("Day 2 Results")
	task1Result := task1.Task1()
	fmt.Println("Task1 result:", task1Result) // 407
	task2Result := task2.Task2()
	fmt.Println("Task2 result:", task2Result) // 459
}
