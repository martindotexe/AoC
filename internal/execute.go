package internal

import (
	"fmt"
	"os"

	"martindotexe/AoC/2024/puzzles/day01"
	"martindotexe/AoC/2024/puzzles/day02"
	"martindotexe/AoC/2024/puzzles/day03"
	"martindotexe/AoC/2024/puzzles/day04"
	"martindotexe/AoC/2024/puzzles/day05"
)

var mappings = map[int]func() (int, int){
	1: day01.Run,
	2: day02.Run,
	3: day03.Run,
	4: day04.Run,
	5: day05.Run,
}

func Run(day int) {
	var output string
	if day == 0 {
		output = runAll()
	} else {
		output = runDay(day)
	}
	if output == "" {
		os.Exit(1)
		return
	}
	fmt.Println(output)
}

func runDay(day int) string {
	fn, ok := mappings[day]
	if !ok {
		return ""
	}
	part1, part2 := fn()
	return format(day, part1, part2)
}

func runAll() string {
	output := ""
	for day, mapping := range mappings {
		part1, part2 := mapping()
		output += format(day, part1, part2)
	}
	return output
}

func format(day, part1, part2 int) string {
	return fmt.Sprintf("Day %d:\n\tPart 1: %d\n\tPart 2: %d\n", day, part1, part2)
}
