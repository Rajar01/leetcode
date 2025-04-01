package main

func RemoveDuplicates(nums []int) int {
	slow := 0
	fast := slow + 1

	for slow < len(nums)-1 {
		if nums[slow] == nums[fast] {
			nums = append(nums[:fast], nums[fast+1:]...)
			fast--
		}

		if fast == len(nums)-1 {
			slow++
			fast = slow + 1
		} else {
			fast++
		}
	}

	return len(nums)
}
