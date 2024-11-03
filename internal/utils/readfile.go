package utils

import (
	"bufio"
	"iter"
	"os"
	"path/filepath"
)

func IterLines(path string) iter.Seq[string] {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(pwd, path)
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
