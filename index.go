package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	left := make([]int, 0)
	right := make([]int, 0)

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile("\\s+")

	lineCount := 0

	for scanner.Scan() {
		lineCount++
		result := regex.Split(scanner.Text(), 2)
		leftToNumber, err := strconv.Atoi(result[0])
		if err != nil {
			panic(err)
		}
		rightToNumber, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}
		left = append(left, leftToNumber)
		right = append(right, rightToNumber)
	}

	fmt.Println("Line count:", lineCount)

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	totalDistance := 0
	for i := 0; i < lineCount; i++ {
		totalDistance += abs(left[i] - right[i])
	}

	fmt.Println("Total distance:", totalDistance)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
