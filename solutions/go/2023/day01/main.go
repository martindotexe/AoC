package main

import (
	"fmt"
	"iter"
	"os"
	"slices"
	"strings"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
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
	sums := 0
	for line := range in {
		l, r := 0, len(line)-1
		for l <= r {
			if !isDigit(line[l]) {
				l++
			}
			if !isDigit(line[r]) {
				r--
			}
			if isDigit(line[l]) && isDigit(line[r]) {
				break
			}
		}
		sums += toInt(line[l])*10 + toInt(line[r])
	}
	return sums
}

func part2(in iter.Seq[string]) int {
	sums := 0
	for line := range in {
		l, r := 0, len(line)-1
		ls, rs := 0, 0

		for l <= r {
			if d, ok := isSpelledDigit(line[l:]); !ok {
				l++
			} else {
				ls = d
			}
			if d, ok := isSpelledDigit(line[r:]); !ok {
				r--
			} else {
				rs = d
			}
			if ls != 0 && rs != 0 {
				break
			}
		}
		sums += ls*10 + rs
	}
	return sums
}

func toInt(b byte) int {
	return int(b - '0')
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func isSpelledDigit(s string) (int, bool) {
	if isDigit(s[0]) {
		return toInt(s[0]), true
	}
	for spelled, d := range digitMap {
		if strings.HasPrefix(s, spelled) {
			return d, true
		}
	}
	return 0, false
}
