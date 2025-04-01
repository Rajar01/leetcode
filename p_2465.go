package main

import (
	"slices"
	"sort"
)

func DistinctAverages(nums []int) int {
	result := 0
	l := 0
	r := len(nums) - 1
	avg := []float64{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)/2; i++ {
		if !slices.Contains(avg, (float64(nums[l])+float64(nums[r]))/2) {
			avg = append(avg, (float64(nums[l])+float64(nums[r]))/2)
			result++
		}

		l++
		r--
	}

	return result
}
