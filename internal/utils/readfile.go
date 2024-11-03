package utils

import (
	"bufio"
	"iter"
	"os"
)

func IterLines(path string) iter.Seq[string] {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	return func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}
