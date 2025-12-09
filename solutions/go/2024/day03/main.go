package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	input := string(data)

	result1 := part1(input)
	result2 := part2(input)

	fmt.Printf("Part 1: %d\n", result1)
	fmt.Printf("Part 2: %d\n", result2)
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
