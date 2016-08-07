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
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	leftDepth := getDepth(root.Left)
	rightDepth := getDepth(root.Right)
	if math.Abs(float64((leftDepth - rightDepth))) <= 1 {
		return true && isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}
}
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return int(math.Max(float64(getDepth(root.Left)), float64(getDepth(root.Right))) + 1)
}
