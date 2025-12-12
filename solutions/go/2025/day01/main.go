package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

const MAX = 100

func partOne(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var dial, count = 50, 0
	for {
		var turn = 0
		var dir byte
		if _, err := fmt.Fscanf(reader, "%c%d\n", &dir, &turn); err != nil {
			if !errors.Is(err, io.EOF) {
				panic(err)
			}
			break
		}
		if dir == 'L' {
			dial = (dial - turn%MAX + MAX) % MAX
		} else {
			dial = (dial + turn) % MAX
		}
		if dial == 0 {
			count++
		}

	}
	fmt.Println(count)
}

func partTwo(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var dial, count = 50, 0
	for {
		var turn = 0
		var dir byte
		if _, err := fmt.Fscanf(reader, "%c%d\n", &dir, &turn); err != nil {
			if !errors.Is(err, io.EOF) {
				panic(err)
			}
			break
		}

		// Calculate how many steps until we first hit 0, and update dial position
		var firstHit = 0
		if dir == 'L' {
			// Going left (counter-clockwise): we hit 0 after 'dial' steps
			firstHit = dial
			dial = (dial - turn%MAX + MAX) % MAX
		} else {
			// Going right (clockwise): we hit 0 after (MAX - dial) steps
			firstHit = MAX - dial
			dial = (dial + turn) % MAX
		}

		// Handle edge case: if already at 0, first hit is after a full rotation
		if firstHit == 0 {
			firstHit = MAX
		}

		// If we reach the first hit within 'turn' steps, calculate total hits
		// After the first hit, we hit 0 again every MAX steps
		if firstHit <= turn {
			count += 1 + (turn-firstHit)/MAX
		}
	}
	fmt.Println(count)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: go run main.go <filepath>\n")
		os.Exit(1)
	}
	var path = os.Args[1]
	partOne(path)
	partTwo(path)
}
