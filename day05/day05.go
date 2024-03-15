package day05

import (
	"aoc2021/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x int
	y int
}

func parseToGrid(input []string, diagonals bool) map[point]int {
	vents := make(map[point]int)
	for _, line := range input {
		s1 := strings.Split(line, " -> ")
		p1, _ := utils.Map(strings.Split(s1[0], ","), strconv.Atoi)
		p2, _ := utils.Map(strings.Split(s1[1], ","), strconv.Atoi)
		from, to := point{p1[0], p1[1]}, point{p2[0], p2[1]}

		if from.x == to.x {
			abs := math.Abs(float64(from.y - to.y))
			yDir := from.y < to.y
			for i := 0; i <= int(abs); i++ {
				var newY int
				if yDir {
					newY = from.y + i
				} else {
					newY = from.y - i
				}
				new := point{from.x, newY}
				vents[new]++
			}
		} else if from.y == to.y {
			abs := math.Abs(float64(from.x - to.x))
			xDir := from.x < to.x
			for i := 0; i <= int(abs); i++ {
				var newX int
				if xDir {
					newX = from.x + i
				} else {
					newX = from.x - i
				}
				new := point{newX, from.y}
				vents[new]++
			}
		} else if diagonals {
			abs := math.Abs(float64(from.x - to.x))
			xDir := from.x < to.x
			yDir := from.y < to.y
			for i := 0; i <= int(abs); i++ {
				var newX int
				var newY int
				if xDir {
					newX = from.x + i
				} else {
					newX = from.x - i
				}
				if yDir {
					newY = from.y + i
				} else {
					newY = from.y - i
				}
				new := point{newX, newY}
				vents[new]++
			}
		}
	}

	return vents
}

func solveOne(input []string) int {
	vents := parseToGrid(input, false)
	acc := 0
	for _, val := range vents {
		if val > 1 {
			acc++
		}
	}
	return acc
}

func solveTwo(input []string) int {
	vents := parseToGrid(input, true)
	acc := 0
	for _, val := range vents {
		if val > 1 {
			acc++
		}
	}
	return acc
}

func Solve() {
	lines, _ := utils.ReadToLines("inputs/day05.txt")

	start := time.Now()
	fmt.Println("Part 1: ", solveOne(lines))
	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
	start = time.Now()
	fmt.Println("Part 2: ", solveTwo(lines))
	elapsed = time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
}
