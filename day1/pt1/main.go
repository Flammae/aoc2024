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

func scan(fileName string, separator *regexp.Regexp, left *[]int, right *[]int) (error, int) {
	file, err := os.Open(fileName)
	if err != nil {
		return err, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0

	for scanner.Scan() {
		lineCount++

		result := separator.Split(scanner.Text(), 2)
		leftToNumber, err := strconv.Atoi(result[0])

		if err != nil {
			return err, lineCount
		}
		rightToNumber, err := strconv.Atoi(result[1])
		if err != nil {
			return err, lineCount
		}
		*left = append(*left, leftToNumber)
		*right = append(*right, rightToNumber)
	}

	return nil, lineCount
}

func sortAscending(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func main() {
	left := make([]int, 0)
	right := make([]int, 0)

	err, lineCount := scan("./input.txt", regexp.MustCompile("\\s+"), &left, &right)

	if err != nil {
		panic(err)
	}

	fmt.Println("Line count:", lineCount)

	sortAscending(left)

	sortAscending(right)

	totalDistance := 0
	for i := 0; i < lineCount; i++ {
		totalDistance += abs(left[i] - right[i])
	}

	fmt.Println("Total distance:", totalDistance)

}
