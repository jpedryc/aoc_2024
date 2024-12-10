package main

import (
	"day9/task1"
	"day9/task2"
	"fmt"
)

func main() {
	fmt.Println("Day 9 Results")
	task1Result := task1.Task1()
	fmt.Println("Task1 result:", task1Result) // 6200294120911
	task2Result := task2.Task2()
	fmt.Println("Task2 result:", task2Result) // 6227018762750
}
