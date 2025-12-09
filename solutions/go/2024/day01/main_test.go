package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := slices.Values([]string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	})

	assert.Equal(t, 11, part1(input))
}

func TestPart2(t *testing.T) {
	input := slices.Values([]string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	})

	assert.Equal(t, 31, part2(input))
}
