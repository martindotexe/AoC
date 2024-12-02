package day01

import (
	"iter"
	"slices"
	"strconv"
	"strings"

	"martindotexe/AoC/2024/internal/utils"
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

func (n *Node) iter() iter.Seq[*Node] {
	return func(yield func(*Node) bool) {
		if n.left != nil {
			n.left.iter()(yield)
		}
		if !yield(n) {
			return
		}
		if n.right != nil {
			n.right.iter()(yield)
		}
	}
}

func zip(left, right iter.Seq[*Node]) iter.Seq2[*Node, *Node] {
	leftS := slices.Collect(left)
	rightS := slices.Collect(right)
	return func(yield func(*Node, *Node) bool) {
		for i := 0; i < len(leftS); i++ {
			if !yield(leftS[i], rightS[i]) {
				return
			}
		}
	}
}

func Run() (int, int) {
	input1 := utils.IterLines("puzzles/day01/in.txt")
	input2 := utils.IterLines("puzzles/day01/in.txt")
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
	for l, r := range zip(left.iter(), right.iter()) {
		distance := max(l.value, r.value) - min(l.value, r.value)
		sum += distance
	}
	return sum
}

func part2(input iter.Seq[string]) int {
	left := &Node{}
	right := map[int]int{}
	for i, line := range utils.Enumerate(input) {
		data := strings.Split(line, "   ")
		l, lerr := strconv.Atoi(data[0])
		r, rerr := strconv.Atoi(data[1])
		if lerr != nil || rerr != nil {
			panic("Failed to convert string to int")
		}
		if i == 0 {
			left = newNode(l)
		} else {
			left.push(l)
		}
		right[r] += 1
	}
	sum := 0
	for l := range left.iter() {
		simScore := right[l.value]
		sum += l.value * simScore
	}
	return sum
}
