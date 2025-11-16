package day05

import (
	"sort"
	"strconv"
	"strings"

	"martindotexe/AoC/internal/utils"
)

func Run() (int, int) {
	input := utils.ReadFile("../data/2024/day05.txt")
	return part1(input), part2(input)
}

func parseRules(input []string) (map[[2]int]bool, int) {
	rules := make(map[[2]int]bool)
	var i int
	for i = 0; i < len(input) && input[i] != ""; i++ {
		line := input[i]
		nums := strings.Split(line, "|")

		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		rules[[2]int{a, b}] = true
		rules[[2]int{b, a}] = false
	}
	return rules, i + 1
}

type sortByRule struct {
	rules map[[2]int]bool
	objs  []int
}

func (a sortByRule) Len() int           { return len(a.objs) }
func (a sortByRule) Swap(i, j int)      { a.objs[i], a.objs[j] = a.objs[j], a.objs[i] }
func (a sortByRule) Less(i, j int) bool { return a.rules[[2]int{a.objs[i], a.objs[j]}] }

func part1(input []string) int {
	rules, i := parseRules(input)

	sum := 0

	for ; i < len(input); i++ {
		nums := strings.Split(input[i], ",")
		sortable := &sortByRule{rules, utils.ToInts(nums)}
		if sort.IsSorted(sortable) {
			sum += sortable.objs[len(sortable.objs)/2]
		}
	}
	return sum
}

func part2(input []string) int {
	rules, i := parseRules(input)

	sum := 0

	for ; i < len(input); i++ {
		nums := strings.Split(input[i], ",")
		sortable := &sortByRule{rules, utils.ToInts(nums)}
		if !sort.IsSorted(sortable) {
			sort.Sort(sortable)
			sum += sortable.objs[len(sortable.objs)/2]
		}
	}
	return sum
}
