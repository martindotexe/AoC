package day01

import (
	"iter"
	"strings"

	"martindotexe/AoC/internal/utils"
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

func Run() (int, int) {
	in1 := utils.IterLines("../data/2023/day01.txt")
	in2 := utils.IterLines("../data/2023/day01.txt")

	return Part1(in1), Part2(in2)
}

func Part1(in iter.Seq[string]) int {
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

func Part2(in iter.Seq[string]) int {
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
