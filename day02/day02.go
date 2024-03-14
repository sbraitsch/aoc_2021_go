package day02

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solveOne(movements []string) int {
	start := time.Now()
	x, y := 0, 0
	for _, movement := range movements {
		split := strings.Split(movement, " ")
		length, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			x += length
		case "up":
			y -= length
		case "down":
			y += length
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
	return x * y
}

func solveTwo(movements []string) int {
	start := time.Now()
	x, y, aim := 0, 0, 0
	for _, movement := range movements {
		split := strings.Split(movement, " ")
		length, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			x += length
			y += length * aim
		case "up":
			aim -= length
		case "down":
			aim += length
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/1_000_000)
	return x * y
}

func Solve() {
	lines, _ := utils.ReadToLines("inputs/day02.txt")
	fmt.Println("Part 1: ", solveOne(lines))
	fmt.Println("Part 2: ", solveTwo(lines))
}
