package binaryTree

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	ret := []int{}
	for len(queue) > 0 {
		arrays := []*TreeNode{}
		i := 0
		for ; i < len(queue); i++ {
			tmp := queue[i]
			if tmp.LChild != nil {
				arrays = append(arrays, tmp.LChild)
			}
			if tmp.RChild != nil {
				arrays = append(arrays, tmp.RChild)
			}
		}
		ret = append(ret, queue[i-1].Val)
		queue = arrays
	}
	return ret
}
