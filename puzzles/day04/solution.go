package day04

import (
	"iter"
	"strings"

	"martindotexe/aoc/internal/utils"
)

func Run() (int, int) {
	in := utils.IterLines("puzzles/day04/in.txt")

	return Part1(in), Part2()
}

func Part1(in iter.Seq[string]) int {
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

func Part2() int {
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
