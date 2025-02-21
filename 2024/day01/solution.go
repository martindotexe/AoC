package day01

import (
	"iter"
	"strconv"
	"strings"

	"martindotexe/AoC/internal/utils"
)

type Node struct {
	left  *Node
	right *Node
	value int
}

func (n *Node) push(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.push(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.push(value)
		}
	}
}

func newNode(value int) *Node {
	return &Node{value: value}
}

func (n *Node) slice() []int {
	s := []int{}
	var f func(*Node)
	f = func(n *Node) {
		if n.left != nil {
			f(n.left)
		}
		s = append(s, n.value)
		if n.right != nil {
			f(n.right)
		}
	}
	f(n)
	return s
}

func zip(left, right []int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i := 0; i < len(left); i++ {
			if !yield(left[i], right[i]) {
				return
			}
		}
	}
}

func Run() (int, int) {
	input1 := utils.IterLines("2024/day01/in.txt")
	input2 := utils.IterLines("2024/day01/in.txt")
	return part1(input1), part2(input2)
}

func part1(input iter.Seq[string]) int {
	left, right := &Node{}, &Node{}
	for i, line := range utils.Enumerate(input) {
		data := strings.Split(line, "   ")
		l, lerr := strconv.Atoi(data[0])
		r, rerr := strconv.Atoi(data[1])
		if lerr != nil || rerr != nil {
			panic("Failed to convert string to int")
		}
		if i == 0 {
			left = newNode(l)
			right = newNode(r)
		} else {
			left.push(l)
			right.push(r)
		}
	}
	sum := 0
	for l, r := range zip(left.slice(), right.slice()) {
		distance := max(l, r) - min(l, r)
		sum += distance
	}
	return sum
}

func part2(input iter.Seq[string]) int {
	left := []int{}
	right := map[int]int{}
	for line := range input {
		data := strings.Split(line, "   ")
		l, lerr := strconv.Atoi(data[0])
		r, rerr := strconv.Atoi(data[1])
		if lerr != nil || rerr != nil {
			panic("Failed to convert string to int")
		}
		left = append(left, l)
		right[r] += 1
	}
	sum := 0
	for _, l := range left {
		simScore := right[l]
		sum += l * simScore
	}
	return sum
}
