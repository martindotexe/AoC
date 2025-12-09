package main

import (
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := slices.Values([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	})

	assert.Equal(t, 142, part1(input))
}

func TestPart2(t *testing.T) {
	input := slices.Values([]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	})

	assert.Equal(t, 281, part2(input))
}

func BenchmarkPart1(b *testing.B) {
	data, _ := os.ReadFile("../../../data/2023/day01.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	in := slices.Values(lines)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(in)
	}
}

func BenchmarkPart2(b *testing.B) {
	data, _ := os.ReadFile("../../../data/2023/day01.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	in := slices.Values(lines)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(in)
	}
}
