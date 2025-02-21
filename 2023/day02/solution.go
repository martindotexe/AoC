package day02

import (
	"iter"
	"strconv"
	"strings"

	"martindotexe/AoC/internal/utils"
)

var maxColors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Run() (int, int) {
	in1 := utils.IterLines("2023/day02/in.txt")
	in2 := utils.IterLines("2023/day02/in.txt")

	return Part1(in1), Part2(in2)
}

func Part1(in iter.Seq[string]) int {
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

func Part2(in iter.Seq[string]) int {
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
