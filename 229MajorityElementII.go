package main

import (
	"fmt"
)

func majorityElement(nums []int) []int {
	var n1, n2 int
	c1 := 0
	c2 := 0
	for i := range nums {
		if n1 == nums[i] {
			c1++
		} else if n2 == nums[i] {
			c2++
		} else if c1 == 0 {
			n1 = nums[i]
			c1++
		} else if c2 == 0 {
			n2 = nums[i]
			c2++
		} else {
			c1--
			c2--
		}
	}
	c1 = 0
	c2 = 0
	for i := range nums {
		if n1 == nums[i] {
			c1++
		} else if n2 == nums[i] {
			c2++
		}
	}
	var result []int
	if c1 > len(nums)/3 {
		result = append(result, n1)
	}

	if c2 > len(nums)/3 {
		result = append(result, n2)
	}
	return result
}
