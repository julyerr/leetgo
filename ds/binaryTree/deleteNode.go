package binaryTree

// using recursion to handle parent,child problems
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.LChild == nil && root.RChild == nil {
			return nil
		}
		if root.RChild == nil {
			root.Val = root.LChild.Val
			root.RChild = root.LChild.RChild
			root.LChild = root.LChild.LChild
		} else {
			if root.RChild.LChild == nil {
				root.Val = root.RChild.Val
				root.RChild = root.RChild.RChild
			} else {
				head := root
				node := root.RChild
				for node.LChild != nil {
					head = node
					node = node.LChild
				}
				root.Val = node.Val
				head.LChild = node.RChild
			}
		}
	} else if root.Val > key {
		if root.LChild == nil {
			return root
		}
		root.LChild = DeleteNode(root.LChild, key)
	} else {
		if root.RChild == nil {
			return root
		}
		root.RChild = DeleteNode(root.RChild, key)
	}
	return root
}
