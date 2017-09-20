package binaryTree

import (
	"math"
)

func IsValidBST(root *TreeNode) bool {
	p := root
	stack := []*TreeNode{}
	ret := []int{}
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.LChild
		}
		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = tmp.RChild
			if len(ret) > 0 && tmp.Val <= ret[len(ret)-1] {
				return false
			}
			ret = append(ret, tmp.Val)
		}
	}
	return true
}

func IsValidBST1(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return traversal(root, math.MinInt32, math.MaxInt32)
}

// just pay attention to the sort func , rather than using inorder traversal
func traversal(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val < min || root.Val > max {
		return false
	}
	return traversal(root.LChild, min, root.Val-1) && traversal(root.RChild, root.Val+1, max)
}
