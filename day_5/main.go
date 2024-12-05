package main

import (
	"day5/task1"
	"day5/task2"
	"fmt"
)

func main() {
	fmt.Println("Day 5 Results")
	task1Result := task1.Task1()
	fmt.Println("Task1 result:", task1Result) // 5091
	task2Result := task2.Task2()
	fmt.Println("Task2 result:", task2Result) // 4681
}
