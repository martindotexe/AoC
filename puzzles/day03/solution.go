package day03

import (
	"os"
	"regexp"
	"strconv"
)

func Run() (int, int) {
	data, err := os.ReadFile("puzzles/day03/in.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	return part1(input), part2(input)
}

func part1(input string) int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		a, aerr := strconv.Atoi(match[1])
		b, berr := strconv.Atoi(match[2])
		if aerr != nil || berr != nil {
			panic("Failed to convert string to int")
		}
		sum += a * b
	}
	return sum
}

func part2(input string) int {
	regex := regexp.MustCompile(`(mul|do|don't)\((?:(\d+),(\d+)|)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)

	sum := 0
	enabled := true
	for _, match := range matches {
		if match[1] == "do" {
			enabled = true
		} else if match[1] == "don't" {
			enabled = false
		} else if enabled && match[1] == "mul" {
			a, aerr := strconv.Atoi(match[2])
			b, berr := strconv.Atoi(match[3])
			if aerr != nil || berr != nil {
				panic("Failed to convert string to int")
			}
			sum += a * b
		}
	}
	return sum
}
