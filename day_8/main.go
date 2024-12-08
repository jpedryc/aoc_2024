package main

import (
	"day8/task1"
	"day8/task2"
	"fmt"
)

func main() {
	fmt.Println("Day 8 Results")
	task1Result := task1.Task1()
	fmt.Println("Task1 result:", task1Result) // 341
	task2Result := task2.Task2()
	fmt.Println("Task2 result:", task2Result) // 1134
}
