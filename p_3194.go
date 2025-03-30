package main

import (
	"sort"
)

func MinimumAverage(nums []int) float64 {
	var result float64
	n := len(nums) / 2
	l := 0
	r := len(nums) - 1

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < n; i++ {
		if float64((nums[l]+nums[r])/2) < result || i == 0 {
			result = (float64(nums[l]) + float64(nums[r])) / 2
		}

		l++
		r--
	}

	return result
}
