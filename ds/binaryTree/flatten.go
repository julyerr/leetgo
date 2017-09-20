package binaryTree

// 做题不仅是锻炼自己的思维方式，更加重要的是常见的问题的解决方法
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var pre *TreeNode
	flatten(root, &pre)
}

func flatten(root *TreeNode, prev **TreeNode) {
	if *prev != nil {
		// values & variable pointer diffs
		(*prev).LChild = nil
		(*prev).RChild = root
	}
	*prev = root
	// 间接修改指针的内容，导致超时、出错
	// if root.LChild != nil {
	// 	flatten(root.LChild, prev)
	// }
	// if root.RChild != nil {
	// 	flatten(root.RChild, prev)
	// }

	left := root.LChild
	right := root.RChild
	if left != nil {
		flatten(left, prev)
	}
	if right != nil {
		flatten(right, prev)
	}
}
