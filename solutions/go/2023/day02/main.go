package main

import (
	"fmt"
	"iter"
	"os"
	"slices"
	"strconv"
	"strings"
)

var maxColors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

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
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	in1 := slices.Values(lines)
	in2 := slices.Values(lines)

	result1 := part1(in1)
	result2 := part2(in2)

	fmt.Printf("Part 1: %d\n", result1)
	fmt.Printf("Part 2: %d\n", result2)
}

func part1(in iter.Seq[string]) int {
	sum := 0
	for line := range in {
		game := strings.Split(line, ": ")
		rounds := strings.Split(game[1], "; ")
		gameId, err := strconv.Atoi(strings.Split(game[0], " ")[1])
		if err != nil {
			panic(err)
		}

		possible := func() bool {
			for _, round := range rounds {
				hands := strings.Split(round, ", ")
				for _, hand := range hands {
					h := strings.Split(hand, " ")
					color := h[1]
					digit, err := strconv.Atoi(h[0])
					if err != nil {
						panic(err)
					}
					if maxColor, ok := maxColors[color]; !ok {
						panic("Invalid color")
					} else if digit > maxColor {
						return false
					}
				}
			}
			return true
		}()
		if possible {
			sum += gameId
		}
	}
	return sum
}

func part2(in iter.Seq[string]) int {
	sum := 0
	for line := range in {
		game := strings.Split(line, ": ")
		rounds := strings.Split(game[1], "; ")

		minColors := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range rounds {
			hands := strings.Split(round, ", ")
			for _, hand := range hands {
				h := strings.Split(hand, " ")
				color := h[1]
				digit, err := strconv.Atoi(h[0])
				if err != nil {
					panic(err)
				}
				minColors[color] = max(digit, minColors[color])
			}
		}
		sum += minColors["red"] * minColors["blue"] * minColors["green"]
	}
	return sum
}
