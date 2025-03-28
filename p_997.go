package main

import (
	"slices"
)

func FindJudge(n int, trust [][]int) int {
	var result int
	al := map[int][]int{}

	for i := 1; i <= n; i++ {
		al[i] = []int{}
	}

	for _, v := range trust {
		al[v[0]] = append(al[v[0]], v[1])
	}

	for k, v := range al {
		if len(v) == 0 {
			result = k
		}
	}

	for k, v := range al {
		if !slices.Contains(v, result) && k != result {
			return -1
		}
	}

	return result
}
