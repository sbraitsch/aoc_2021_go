package main

import (
	"aoc2021/utils"
	"fmt"
	"time"
)

func solveOne(input []string) int {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
	return 0
}

func solveTwo(input []string) int {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/1_000_000)
	return 0
}

func Solve() {
	lines, _ := utils.ReadToLines("inputs/day02.txt")
	fmt.Println("Part 1: ", solveOne(lines))
	fmt.Println("Part 2: ", solveTwo(lines))
}
