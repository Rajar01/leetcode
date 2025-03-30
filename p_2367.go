package main

import (
	"sort"
)

func ArithmeticTriplets(nums []int, diff int) int {
	result := 0
	n := len(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[k]-nums[j] == diff && nums[j]-nums[i] == diff {
					result++
				}
			}
		}
	}

	return result
}
