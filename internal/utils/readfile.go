package utils

import (
	"bufio"
	"iter"
	"os"
	"path/filepath"
)

func openFile(path string) *os.File {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(pwd, path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func IterLines(path string) iter.Seq[string] {
	file := openFile(path)

	scanner := bufio.NewScanner(file)
	return func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				file.Close()
				return
			}
		}
	}
}

func ReadFile(path string) []string {
	file := openFile(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
