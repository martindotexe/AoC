package day04

import (
	"iter"
	"math"
	"strings"

	"martindotexe/aoc/internal/utils"
)

func Run() (int, int) {
	in := utils.IterLines("puzzles/day04/in.txt")

	return Part1(in), Part2()
}

func Part1(in iter.Seq[string]) int {
	sum := 0
	for line := range in {
		winning, numbers := parse(line)
		winningMap := map[int]bool{}
		for _, win := range winning {
			winningMap[win] = true
		}

		winningTickets := 0
		for _, number := range numbers {
			if winningMap[number] {
				winningTickets++
			}
		}
		sum += int(math.Pow(2, float64(winningTickets-1)))

	}
	return sum
}

func Part2() int {
	return 0
}

func parse(s string) ([]int, []int) {
	s = strings.Split(s, ": ")[1]
	parts := strings.Split(s, "|")
	out := [][]int{}
	for _, part := range parts {
		numbers := []int{}
		currentNumber := 0
		for p := range part {
			if !isDigit(part[p]) {
				continue
			}
			currentNumber = currentNumber*10 + int(part[p]-'0')
			if p+1 >= len(part) || !isDigit(part[p+1]) {
				numbers = append(numbers, currentNumber)
				currentNumber = 0
			}
		}
		out = append(out, numbers)
	}
	return out[0], out[1]
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
