package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	for i := range nums {
		if target == nums[i] {
			return i
		}
	}

	if target < nums[0] {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	for i := range nums {
		if nums[i] < target && target < nums[i+1] {
			return i + 1
		}
	}
	return 0
}
