package day02

import (
	"slices"
	"testing"

	"martindotexe/AoC/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	in := slices.Values([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	})

	assert.Equal(t, 8, Part1(in))
}

func TestPart2(t *testing.T) {
	in := slices.Values([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	})

	assert.Equal(t, 2286, Part2(in))
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
