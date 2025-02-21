package day01

import (
	"slices"
	"testing"

	"martindotexe/AoC/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := slices.Values([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	})

	assert.Equal(t, 142, Part1(input))
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

	assert.Equal(t, 281, Part2(input))
}

func BenchmarkPart1(b *testing.B) {
	in := utils.IterLines("in.txt")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part1(in)
	}
}

func BenchmarkPart2(b *testing.B) {
	in := utils.IterLines("in.txt")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part2(in)
	}
}
