package main

import (
	"day3/task1"
	"day3/task2"
	"fmt"
)

func main() {
	fmt.Println("Day 3 Results")
	task1Result := task1.Task1()
	fmt.Println("Task1 result:", task1Result) // 183788984
	task2Result := task2.Task2()
	fmt.Println("Task2 result:", task2Result) // 62098619
}
