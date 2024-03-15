package main

import (
	"aoc2021/utils"
	"fmt"
	"time"
)

func solveOne(input []string) int {
	return 0
}

func solveTwo(input []string) int {
	return 0
}

func Solve() {
	lines, _ := utils.ReadToLines("inputs/day02.txt")

	start := time.Now()
	fmt.Println("Part 1: ", solveOne(lines))
	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
	start = time.Now()
	fmt.Println("Part 2: ", solveTwo(lines))
	elapsed = time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
}
