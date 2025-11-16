package day02

import (
	"slices"
	"strings"

	"martindotexe/AoC/internal/utils"
)

func Run() (int, int) {
	input := utils.ReadFile("../data/2024/day02.txt")
	return part1(input), part2(input)
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

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		nums := utils.ToInts(strings.Split(line, " "))
		if valid(nums) {
			sum++
		}
	}
	return sum
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		nums := utils.ToInts(strings.Split(line, " "))
		for i := 0; i < len(nums); i++ {
			if valid(slices.Concat(nums[:i], nums[i+1:])) {
				sum++
				break
			}
		}
	}
	return sum
}
