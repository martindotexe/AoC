package day05

import (
	"strconv"
	"strings"

	"martindotexe/AoC/2024/internal/utils"
)

func Run() (int, int) {
	input := utils.ReadFile("puzzles/day05/in.txt")
	return part1(input), part2(input)
}

func checkPages(pages []int, rules map[int][]int) bool {
	illegal := map[int]bool{}

	for _, page := range pages {
		if illegal[page] {
			return false
		}
		if rule, ok := rules[page]; ok {
			for _, r := range rule {
				illegal[r] = true
			}
		}
	}

	return true
}

func parseRules(input []string) (map[int][]int, int) {
	rules := map[int][]int{}
	var i int
	for i = 0; i < len(input); i++ {
		if input[i] == "" {
			break
		}
		line := input[i]
		nums := strings.Split(line, "|")

		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		rules[b] = append(rules[b], a)
	}
	return rules, i
}

func part1(input []string) int {
	rules, i := parseRules(input)
	sum := 0

	for j := i + 1; j < len(input); j++ {
		p := strings.Split(input[j], ",")
		pages := utils.ToInts(p)

		if checkPages(pages, rules) {
			sum += pages[len(pages)/2]
		}

	}

	return sum
}

type Node struct {
	left  *Node
	right *Node
	value int
}

func (n *Node) push(rules map[int][]int, value int) {
	for _, rule := range rules[value] {
		if rule == n.value {
			if n.left == nil {
				n.left = newNode(value)
			} else {
				n.left.push(rules, value)
			}
			return
		}
	}
	if n.right == nil {
		n.right = newNode(value)
	} else {
		n.right.push(rules, value)
	}
}

func newNode(value int) *Node {
	return &Node{value: value}
}

func (n *Node) slice() []int {
	s := []int{}
	if n.left != nil {
		s = append(s, n.left.slice()...)
	}
	s = append(s, n.value)
	if n.right != nil {
		s = append(s, n.right.slice()...)
	}
	return s
}

func part2(input []string) int {
	rules, i := parseRules(input)
	sum := 0

	for j := i + 1; j < len(input); j++ {
		p := strings.Split(input[j], ",")
		pages := utils.ToInts(p)

		if !checkPages(pages, rules) {
			t := newNode(pages[0])
			for _, page := range pages[1:] {
				t.push(rules, page)
			}
			sum += t.slice()[len(pages)/2]
		}

	}

	return sum
}
