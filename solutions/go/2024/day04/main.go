package main

import (
	"fmt"
	"os"
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
	input := string(data)

	result1 := part1(input)
	result2 := part2(input)

	fmt.Printf("Part 1: %d\n", result1)
	fmt.Printf("Part 2: %d\n", result2)
}

func parse(input string) [][]rune {
	out := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		l := make([]rune, len(line))
		for i, c := range line {
			l[i] = c
		}
		out = append(out, l)
	}
	return out
}

func isPattern(g [][]rune, r, c, dr, dc int, pattern string) bool {
	// Check if pattern fits in grid in the direction of dr, dc from r, c.
	i := len(pattern) - 1

	if tr, tc := r+i*dr, c+i*dc; tr < 0 || tr >= len(g) || tc < 0 || tc >= len(g[0]) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if g[r+i*dr][c+i*dc] != rune(pattern[i]) {
			return false
		}
	}
	return true
}

func part1(input string) int {
	sum := 0

	g := parse(input)
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			if g[r][c] != 'X' {
				continue
			}
			for _, dr := range []int{-1, 0, 1} {
				for _, dc := range []int{-1, 0, 1} {
					if dr == 0 && dc == 0 {
						continue
					}
					if isPattern(g, r, c, dr, dc, "XMAS") {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func isXPattern(g [][]rune, r, c int, pattern string) bool {
	i := (len(pattern) - 1) / 2
	if r-i < 0 || r+i >= len(g) || c-i < 0 || c+i >= len(g[0]) {
		return false
	}
	if !(isPattern(g, r-i, c-i, 1, 1, pattern) || isPattern(g, r+i, c+i, -1, -1, pattern)) {
		return false
	}
	if !(isPattern(g, r+i, c-i, -1, 1, pattern) || isPattern(g, r-i, c+i, 1, -1, pattern)) {
		return false
	}
	return true
}

func part2(input string) int {
	sum := 0

	g := parse(input)
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			if g[r][c] != 'A' {
				continue
			}
			if isXPattern(g, r, c, "MAS") {
				sum++
			}
		}
	}
	return sum
}
