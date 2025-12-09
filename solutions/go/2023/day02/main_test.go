package main

import (
	"os"
	"slices"
	"strings"
	"testing"

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

	assert.Equal(t, 8, part1(in))
}

func TestPart2(t *testing.T) {
	in := slices.Values([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	})

	assert.Equal(t, 2286, part2(in))
}

func BenchmarkPart1(b *testing.B) {
	data, _ := os.ReadFile("../../../data/2023/day02.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	in := slices.Values(lines)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(in)
	}
}

func BenchmarkPart2(b *testing.B) {
	data, _ := os.ReadFile("../../../data/2023/day02.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	in := slices.Values(lines)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(in)
	}
}
