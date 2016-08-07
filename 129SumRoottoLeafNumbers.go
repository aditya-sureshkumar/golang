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
func sumatLeaf(node *TreeNode, sum *int, num int) {

	if node.Left == nil && node.Right == nil {
		newsum := num*10 + node.Val + *sum
		*sum = newsum
	}
	if node.Left != nil {
		sumatLeaf(node.Left, sum, num*10+node.Val)
	}
	if node.Right != nil {
		sumatLeaf(node.Right, sum, num*10+node.Val)
	}
}
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	sum := 0
	if root.Left != nil {
		sumatLeaf(root.Left, &sum, root.Val)
	}
	if root.Right != nil {
		sumatLeaf(root.Right, &sum, root.Val)
	}
	return sum
}
