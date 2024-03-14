package day04

import (
	"aoc2021/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

type bingo struct {
	fields []int
	won    bool
	size   int
	lines  [][]int
}

func newBingo(fields []int, size int) *bingo {
	b := bingo{fields: fields, size: size}
	b.extractLines()
	return &b
}

func (board *bingo) extractLines() {
	for i := range board.size {
		board.lines = append(board.lines, board.fields[i*5:i*5+5])
		col := make([]int, board.size)
		for j := range board.size {
			col[j] = board.fields[i+j*board.size]
		}
		board.lines = append(board.lines, col)
	}
}

func (board *bingo) checkForBingo(drawn []int) (win bool) {
	for _, line := range board.lines {
		count := 0
		for _, num := range line {
			if !slices.Contains(drawn, num) {
				break
			}
			count++
		}
		if count == 5 {
			board.won = true
			return true
		}
	}

	return false
}

func (board bingo) calculateScore(drawnTillNow []int) int {
	score := 0
	for _, field := range board.fields {
		if !slices.Contains(drawnTillNow, field) {
			score += field
		}
	}
	return score * drawnTillNow[len(drawnTillNow)-1]
}

func solveOne(boards []*bingo, drawn []int) int {
	start := time.Now()
	for i := range drawn {
		for _, board := range boards {
			slice := drawn[:i+1]
			board.checkForBingo(slice)
			if board.won {
				board.won = false
				return board.calculateScore(slice)
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/float64(time.Millisecond))
	return 0
}

func solveTwo(boards []*bingo, drawn []int) int {
	start := time.Now()
	winners := 0
	for i := range drawn {
		for _, board := range boards {
			if !board.won {
				slice := drawn[:i+1]
				board.checkForBingo(slice)
				if board.won {
					winners++
					if winners == len(boards) {
						return board.calculateScore(slice)
					}
				}
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Function execution time: %.2f ms\n", float64(elapsed.Nanoseconds())/1_000_000)
	return 0
}

func Solve() {
	lines, _ := utils.ReadToLines("inputs/day04.txt")
	numSplit := strings.Split(lines[0], ",")
	drawn := make([]int, len(numSplit))
	for i, n := range numSplit {
		num, _ := strconv.Atoi(n)
		drawn[i] = num
	}

	boards := make([]*bingo, len(lines)/6)

	for i := range boards {
		fields := make([]int, 25)
		for j := range 5 {
			index := 2 + (6 * i) + j
			for k, s := range strings.Fields(lines[index]) {
				num, _ := strconv.Atoi(s)
				fields[k+j*5] = num
			}
		}
		boards[i] = newBingo(fields, 5)
	}

	fmt.Println("Part 1: ", solveOne(boards, drawn))
	fmt.Println("Part 2: ", solveTwo(boards, drawn))
}
