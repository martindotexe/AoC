package utils

import "strconv"

func ToInts(in []string) []int {
	out := make([]int, len(in))
	for i, v := range in {
		o, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		out[i] = o
	}
	return out
}
