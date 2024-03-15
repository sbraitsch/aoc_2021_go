package day06

import (
	"aoc2021/utils"
	"fmt"
	"time"
)

func procreate(fishes map[int]int, days int) int {
	for range days {
		newFish := fishes[0]
		for i := range 8 {
			fishes[i] = fishes[i+1]
		}
		fishes[6] += newFish
		fishes[8] = newFish
	}
	acc := 0
	for _, v := range fishes {
		acc += v
	}
	return acc
}

func solveOne(fishes map[int]int) int {
	return procreate(fishes, 80)
}

func solveTwo(fishes map[int]int) int {
	return procreate(fishes, 256)
}

func Solve() {
	numbers, _ := utils.ReadSingleLineToIntSlice("inputs/day06.txt")
	fishes := make(map[int]int)
	for _, timer := range numbers {
		fishes[timer]++
	}

	start := time.Now()
	fmt.Println("Part 1: ", solveOne(fishes))
	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))

	fishes = make(map[int]int)
	for _, timer := range numbers {
		fishes[timer]++
	}

	start = time.Now()
	fmt.Println("Part 2: ", solveTwo(fishes))
	elapsed = time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
}
