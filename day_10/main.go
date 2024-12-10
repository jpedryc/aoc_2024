package main

import (
	"day10/task1"
	"day10/task2"
	"fmt"
)

func main() {
	fmt.Println("Day 10 Results")
	task1Result := task1.Task1()
	fmt.Println("Task1 result:", task1Result) // 472
	task2Result := task2.Task2()
	fmt.Println("Task2 result:", task2Result) // 969
}
