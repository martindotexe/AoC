package main

import (
	"fmt"
	"os"
	"sort"
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
		sortable := &sortByRule{rules, toInts(nums)}
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
		sortable := &sortByRule{rules, toInts(nums)}
		if !sort.IsSorted(sortable) {
			sort.Sort(sortable)
			sum += sortable.objs[len(sortable.objs)/2]
		}
	}
	return sum
}
