package main

import (
	"flag"

	"martindotexe/AoC/internal"
)

func main() {
	d := flag.Int("d", 0, "Choose day to execute")
	y := flag.Int("y", 0, "Choose year to execute")

	flag.Parse()

	internal.Run(*y, *d)
}
