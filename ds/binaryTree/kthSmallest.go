package binaryTree

func KthSmallest(root *TreeNode, k int) int {
	p := root
	stack := []*TreeNode{}
	ret := []int{}
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.LChild
		}
		if len(stack) > 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, tmp.Val)
			p = tmp.RChild
		}
	}
	return ret[k-1]
}
