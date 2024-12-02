package utils

import "iter"

func Enumerate[T any](it iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for j := range it {
			if !yield(i, j) {
				return
			}
			i++
		}
	}
}
