package day03

import "martindotexe/AoC/internal/utils"

func Run() (int, int) {
	in := utils.ReadFile("2024/day03/in.txt")

	return Part1(in), Part2(in)
}

type Point struct {
	x, y int
}

type Part struct {
	char  rune
	point Point
}

var matrix = [8][2]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func Part1(in []string) int {
	partSet := map[Point]bool{}
	parts := make(map[Point][]int, 8) // Preallocate for 8 possible numbers.
	currentNumber := 0

	processNumber := func() {
		if currentNumber != 0 {
			for part := range partSet {
				parts[part] = append(parts[part], currentNumber)
			}
		}
		for part := range partSet {
			delete(partSet, part)
		}
		currentNumber = 0
	}

	for x, line := range in {
		for y, c := range line {
			if isDigit(c) {
				currentNumber = currentNumber*10 + int(c-'0')

				for _, m := range matrix {
					xm, ym := x+m[0], y+m[1]
					if xm < 0 || ym < 0 || xm >= len(in) || ym >= len(in[xm]) {
						continue
					}
					if isPart(rune(in[xm][ym])) {
						partSet[Point{x: xm, y: ym}] = true
					}
				}
			} else {
				processNumber()
			}
		}
		processNumber()
	}

	sum := 0
	for _, numbers := range parts {
		for _, n := range numbers {
			sum += n
		}
	}
	return sum
}

func Part2(in []string) int {
	partSet := map[Part]bool{}
	parts := make(map[Part][]int, 8) // Preallocate for 8 possible numbers.
	currentNumber := 0

	processNumber := func() {
		if currentNumber != 0 {
			for part := range partSet {
				parts[part] = append(parts[part], currentNumber)
			}
		}
		for part := range partSet {
			delete(partSet, part)
		}
		currentNumber = 0
	}

	for x, line := range in {
		for y, c := range line {
			if isDigit(c) {
				currentNumber = currentNumber*10 + int(c-'0')

				for _, m := range matrix {
					xm, ym := x+m[0], y+m[1]
					if xm < 0 || ym < 0 || xm >= len(in) || ym >= len(in[xm]) {
						continue
					}
					if isPart(rune(in[xm][ym])) {
						partSet[Part{char: rune(in[xm][ym]), point: Point{x: xm, y: ym}}] = true
					}
				}
			} else {
				processNumber()
			}
		}
		processNumber()
	}

	sum := 0
	for part, numbers := range parts {
		if part.char == '*' && len(numbers) == 2 {
			sum += numbers[0] * numbers[1]
		}
	}
	return sum
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func isPart(c rune) bool {
	return !isDigit(c) && c != '.'
}
