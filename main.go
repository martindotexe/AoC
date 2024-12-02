package main

import (
	"flag"

	"martindotexe/AoC/2024/internal"
)

func main() {
	d := flag.Int("d", 0, "Choose day to execute")

	flag.Parse()

	internal.Run(*d)
}
