package binaryTree

func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{}
	stack := []*TreeNode{root}
	head := root
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node.LChild == head || node.RChild == head || node.LChild == nil && node.RChild == nil {
			ret = append(ret, node.Val)
			stack = stack[:len(stack)-1]
			head = node
		} else {
			if node.RChild != nil {
				stack = append(stack, node.RChild)
			}
			if node.LChild != nil {
				stack = append(stack, node.LChild)
			}
		}
	}
	return ret
}
