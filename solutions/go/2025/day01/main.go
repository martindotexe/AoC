package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

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
			turn = 100 - turn
		}
		dial = (dial + turn) % 100
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
		var turn, step = 0, 1
		var dir byte
		if _, err := fmt.Fscanf(reader, "%c%d\n", &dir, &turn); err != nil {
			if !errors.Is(err, io.EOF) {
				panic(err)
			}
			break
		}
		if dir == 'L' {
			step = 99
		}
		for range turn {
			dial = (dial + step) % 100
			if dial == 0 {
				count++
			}
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
	// partOne(path)
	partTwo(path)
}
