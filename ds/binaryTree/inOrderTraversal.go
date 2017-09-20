package binaryTree

type TreeNode struct {
	Val    int
	LChild *TreeNode
	RChild *TreeNode
}

func InOrderTraversal(root *TreeNode) []int {
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
			ret = append(ret, tmp.Val)
		}
	}
	return ret
}
