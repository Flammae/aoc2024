package main

import (
	"bufio"
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

type Scanner struct {
	bufioScanner *bufio.Scanner
	separator    *regexp.Regexp
}

func NewScanner(file *os.File, separator *regexp.Regexp) *Scanner {
	return &Scanner{
		bufioScanner: bufio.NewScanner(file),
		separator:    separator,
	}
}

func (s *Scanner) Scan() bool {
	return s.bufioScanner.Scan()
}

func (s *Scanner) Slice() (error, []int) {
	text := s.bufioScanner.Text()
	textSlice := s.separator.Split(text, -1)

	slice := make([]int, len(textSlice))

	for i, str := range textSlice {
		toInt, err := strconv.Atoi(str)
		if err != nil {
			return err, nil
		}
		slice[i] = toInt
	}

	return nil, slice
}

func sortAscending(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := NewScanner(file, regexp.MustCompile(`\s+`))

	safeCount := 0

	for scanner.Scan() {
		err, slice := scanner.Slice()
		if err != nil {
			panic(err)
		}

		isAscending := slice[0] < slice[1]
		isSafe := true
		for i := 1; i < len(slice); i++ {
			prev := slice[i-1]
			curr := slice[i]

			// Unsafe, based on rules
			if prev == curr {
				isSafe = false
				break
			}
			if isAscending && prev > curr {
				isSafe = false
				break
			}
			if !isAscending && prev < curr {
				isSafe = false
				break
			}
			if abs(prev-curr) > 3 {
				isSafe = false
				break
			}
		}

		if isSafe {
			safeCount += 1
		}
	}

	println("Safe count: ", safeCount)
}
