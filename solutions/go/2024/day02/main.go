package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
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
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	result1 := part1(input)
	result2 := part2(input)

	fmt.Printf("Part 1: %d\n", result1)
	fmt.Printf("Part 2: %d\n", result2)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func allAscending(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func allDescending(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			return false
		}
	}
	return true
}

func valid(nums []int) bool {
	isAscending := allAscending(nums)
	isDescending := allDescending(nums)
	for i := 1; i < len(nums); i++ {
		if a := abs(nums[i] - nums[i-1]); !(1 <= a && a <= 3) {
			return false
		}
	}
	return isAscending || isDescending
}

func toInts(in []string) []int {
	out := make([]int, len(in))
	for i, v := range in {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		out[i] = n
	}
	return out
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		nums := toInts(strings.Split(line, " "))
		if valid(nums) {
			sum++
		}
	}
	return sum
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		nums := toInts(strings.Split(line, " "))
		for i := 0; i < len(nums); i++ {
			if valid(slices.Concat(nums[:i], nums[i+1:])) {
				sum++
				break
			}
		}
	}
	return sum
}
