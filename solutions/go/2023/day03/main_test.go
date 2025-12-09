package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	in := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	assert.Equal(t, 4361, part1(in))
}

func TestPart2(t *testing.T) {
	in := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	assert.Equal(t, 467835, part2(in))
}

func BenchmarkPart1(b *testing.B) {
	data, _ := os.ReadFile("../../../data/2023/day03.txt")
	in := strings.Split(strings.TrimSpace(string(data)), "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(in)
	}
}

func BenchmarkPart2(b *testing.B) {
	data, _ := os.ReadFile("../../../data/2023/day03.txt")
	in := strings.Split(strings.TrimSpace(string(data)), "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(in)
	}
}
