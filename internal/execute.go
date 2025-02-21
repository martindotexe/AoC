package internal

import (
	"fmt"
	"os"

	"martindotexe/AoC/2024/day01"
	"martindotexe/AoC/2024/day02"
	"martindotexe/AoC/2024/day03"
	"martindotexe/AoC/2024/day04"
	"martindotexe/AoC/2024/day05"
)

var mappings = map[int]map[int]func() (int, int){
	2024: {
		1: day01.Run,
		2: day02.Run,
		3: day03.Run,
		4: day04.Run,
		5: day05.Run,
	},
}

func Run(year, day int) {
	var output string
	if day == 0 || year == 0 {
		output = runAll()
	} else {
		output = runDay(year, day)
	}
	if output == "" {
		os.Exit(1)
		return
	}
	fmt.Println(output)
}

func runDay(year, day int) string {
	fn, ok := mappings[year][day]
	if !ok {
		return ""
	}
	part1, part2 := fn()
	return format(year, day, part1, part2)
}

func runAll() string {
	output := ""
	for year, yearMapping := range mappings {
		for day, fn := range yearMapping {
			part1, part2 := fn()
			output += format(year, day, part1, part2)
		}
	}
	return output
}

func format(year, day, part1, part2 int) string {
	return fmt.Sprintf("Year %d\n\tDay %d:\n\t\tPart 1: %d\n\t\tPart 2: %d\n", year, day, part1, part2)
}
