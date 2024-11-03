package main

import (
	"flag"

	"martindotexe/aoc/internal"
)

func main() {
	d := flag.Int("d", 0, "Choose day to execute")

	flag.Parse()

	internal.Run(*d)
}
