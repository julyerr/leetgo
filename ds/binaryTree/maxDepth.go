package binaryTree

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hight1 := MaxDepth(root.LChild) + 1
	hight2 := MaxDepth(root.RChild) + 1
	if hight1 > hight2 {
		return hight1
	}
	return hight2
}
