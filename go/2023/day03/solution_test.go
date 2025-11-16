package day03

import (
	"testing"

	"martindotexe/AoC/internal/utils"

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

	assert.Equal(t, 4361, Part1(in))
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

	assert.Equal(t, 467835, Part2(in))
}

func BenchmarkPart1(b *testing.B) {
	in := utils.ReadFile("../../../data/2023/day03.txt")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part1(in)
	}
}

func BenchmarkPart2(b *testing.B) {
	in := utils.ReadFile("../../../data/2023/day03.txt")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part2(in)
	}
}
