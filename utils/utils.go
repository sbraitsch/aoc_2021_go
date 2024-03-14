package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadToLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file", err)
	}

	return lines, nil
}

func ReadToInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file", err)
	}

	return lines, nil
}

func SumSlice(slice []int) int {
	acc := 0
	for _, n := range slice {
		acc += n
	}
	return acc
}

func Map[T any](strings []string, converter func(string) (T, error)) ([]T, error) {
	result := make([]T, len(strings))
	for i, str := range strings {
		target, err := converter(str)
		if err != nil {
			return nil, err
		}
		result[i] = target
	}
	return result, nil
}
