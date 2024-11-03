package day1

import (
	"slices"
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
