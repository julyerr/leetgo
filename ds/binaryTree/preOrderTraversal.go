package binaryTree

func PreOrderTraversal(root *TreeNode) []int {
	p := root
	stack := []*TreeNode{}
	ret := []int{}
	for p != nil || len(stack) != 0 {
		for p != nil {
			ret = append(ret, p.Val)
			stack = append(stack, p)
			p = p.LChild
		}
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.RChild
		}
	}
	return ret
}
