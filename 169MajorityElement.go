package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	candidate := candidateElement(nums)
	count := 0
	for i := range nums {
		if nums[i] == candidate {
			count++
		}
	}
	if count > len(nums)/2 {
		return candidate
	} else {
		return 0
	}
}

func candidateElement(nums []int) int {
	count := 0
	majElement := nums[0]
	for i := range nums {
		if majElement == nums[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majElement = nums[i]
			count++
		}
	}
	return majElement
}
