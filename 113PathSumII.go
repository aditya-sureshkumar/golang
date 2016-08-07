package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	var values [][]int
	arr := []int{}
	if root == nil {
		return values
	}
	hasPathSum(root, sum, arr, &values)
	return values
}
func hasPathSum(root *TreeNode, sum int, a []int, values *[][]int) {
	outarr := make([]int, len(a)+1)
	for i := 0; i < len(a); i++ {
		outarr[i] = a[i]
	}
	fmt.Println(outarr, a)
	outarr[len(a)] = root.Val
	if root.Left == nil && root.Right == nil {
		fmt.Println(outarr)
		sm := 0
		for i := 0; i < len(outarr); i++ {
			sm = sm + outarr[i]
		}
		fmt.Println(sm)
		if sm == sum {
			*values = append(*values, outarr)
		}
		fmt.Println(values)
	}
	if root.Left != nil {
		hasPathSum(root.Left, sum, outarr, values)
	}
	if root.Right != nil {
		hasPathSum(root.Right, sum, outarr, values)
	}
}
