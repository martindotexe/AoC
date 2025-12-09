package main

import (
	"fmt"
	"iter"
	"os"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: go run main.go <filepath>\n")
		os.Exit(1)
	}

	filepath := os.Args[1]
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	in := slices.Values(lines)

	result1 := part1(in)
	result2 := part2()

	fmt.Printf("Part 1: %d\n", result1)
	fmt.Printf("Part 2: %d\n", result2)
}

func part1(in iter.Seq[string]) int {
	sum := 0
	winningMap := make(map[int]bool)
	for line := range in {
		winning, numbers := parse(line)
		for _, win := range winning {
			winningMap[win] = true
		}

		// Start with 1 point, double for each winning number
		points := 1
		for _, number := range numbers {
			if winningMap[number] {
				points <<= 1
			}
		}
		// Right shift to remove the initial 1 point
		sum += points >> 1

		clear(winningMap)

	}
	return sum
}

func part2() int {
	return 0
}

func parse(s string) ([]int, []int) {
	s = strings.Split(s, ": ")[1]
	numbers := []int{}
	winningNumbers := []int{}
	currentNumber := 0

	processNumber := func() {
		if currentNumber != 0 {
			numbers = append(numbers, currentNumber)
		}
		currentNumber = 0
	}

	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			currentNumber = currentNumber*10 + int(s[i]-'0')
		} else if s[i] == '|' {
			winningNumbers = numbers
			numbers = []int{}
		} else {
			processNumber()
		}
	}
	processNumber()
	return winningNumbers, numbers
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
