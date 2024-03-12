package day01

import (
	"aoc2021/utils"
	"fmt"
	"time"
)

func solveOne(measurements []int) int {
	start := time.Now()
	sum := 0
	prev := measurements[0]
	for _, m := range measurements[1:] {
		other := m
		if prev < other {
			sum++
		}
		prev = other
	}
	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
	return sum
}

func solveTwo(measurements []int) int {
	start := time.Now()
	sum := 0
	for i := range len(measurements) {
		cur_window := utils.SumSlice(measurements[i : i+3])
		next_window := utils.SumSlice(measurements[i+1 : i+4])
		if next_window > cur_window {
			sum++
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/1_000_000)
	return sum
}

func Solve() {
	numbers, _ := utils.ReadToInt("inputs/day01.txt")
	fmt.Println("Part 1: ", solveOne(numbers))
	fmt.Println("Part 2: ", solveTwo(numbers))
}
