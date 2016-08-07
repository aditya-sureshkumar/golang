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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if sum-root.Val == 0 {
			return true
		}
		return false
	}
	check := false
	if root.Left != nil {
		check = check || hasPathSum(root.Left, sum-root.Val)
	}
	if root.Right != nil {
		check = check || hasPathSum(root.Right, sum-root.Val)
	}
	return check
}
