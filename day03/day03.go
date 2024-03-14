package day03

import (
	"aoc2021/utils"
	"fmt"
	"math"
	"strconv"
	"time"
)

func filterByCharAt(unfiltered []string, char byte, position int) (ret []string) {
	if len(unfiltered) == 1 {
		return unfiltered
	}
	for _, s := range unfiltered {
		if s[position] == char {
			ret = append(ret, s)
		}
	}
	return
}

func bitWeightAt(binaries []string, position int, inverted bool) byte {
	weight := 0
	for _, binary := range binaries {
		if binary[position] == '0' {
			weight--
		} else {
			weight++
		}
	}

	if (weight >= 0 && inverted) || (weight < 0 && !inverted) {
		return '0'
	} else {
		return '1'
	}
}

func solveOne(bits map[int]int) int {
	start := time.Now()
	maxPos := len(bits) - 1
	gamma, epsilon := 0.0, 0.0
	for pos, bit := range bits {
		if bit > 0 {
			gamma += math.Pow(2, float64(maxPos-pos))
		} else {
			epsilon += math.Pow(2, float64(maxPos-pos))
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
	return int(gamma * epsilon)
}

func solveTwo(binaries []string) int {
	start := time.Now()
	oxySieve, scrubberSieve := binaries, binaries

	for i := range 12 {
		oxySieve = filterByCharAt(oxySieve, bitWeightAt(oxySieve, i, false), i)
		scrubberSieve = filterByCharAt(scrubberSieve, bitWeightAt(scrubberSieve, i, true), i)
	}

	oxy, _ := strconv.ParseInt(oxySieve[0], 2, 64)
	scrubber, _ := strconv.ParseInt(scrubberSieve[0], 2, 64)

	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/1_000_000)
	return int(oxy) * int(scrubber)
}

func Solve() {
	lines, _ := utils.ReadToLines("inputs/day03.txt")
	bitmap := make(map[int]int)
	for _, line := range lines {
		for idx, c := range line {
			switch c {
			case '0':
				bitmap[idx]--
			case '1':
				bitmap[idx]++
			}
		}
	}

	fmt.Println("Part 1: ", solveOne(bitmap))
	fmt.Println("Part 2: ", solveTwo(lines))
}
