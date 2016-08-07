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
func binaryTreePaths(root *TreeNode) []string {
	var values []string
	arr := []int{}
	if root == nil {
		return values
	}
	printPathToRoot(root, arr, &values)
	return values
}
func printPathToRoot(root *TreeNode, a []int, paths *[]string) {
	outarr := make([]int, len(a)+1)
	for i := 0; i < len(a); i++ {
		outarr[i] = a[i]
	}
	//fmt.Println(outarr, a)
	outarr[len(a)] = root.Val

	if root.Left == nil && root.Right == nil {
		var buffer bytes.Buffer
		for i := 0; i < len(outarr); i++ {
			if i != len(outarr)-1 {
				s := strconv.Itoa(outarr[i])
				buffer.WriteString(s)
				buffer.WriteString("->")
			} else {
				s := strconv.Itoa(outarr[i])
				buffer.WriteString(s)
			}
		}
		//fmt.Println(*paths)
		*paths = append(*paths, buffer.String())

	}
	if root.Left != nil {
		printPathToRoot(root.Left, outarr, paths)
	}
	if root.Right != nil {
		printPathToRoot(root.Right, outarr, paths)
	}
}
