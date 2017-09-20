package binaryTree

// when the results will use the previous solutions , considering using recursion
func GenerateTrees(n int) []*TreeNode {
	if n < 1 {
		return nil
	}
	return generateTrees(1, n)
}

// using recursion , consider exit conditions first , then assuming gotten the previous results already , handle the logic
func generateTrees(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}
	if left == right {
		return []*TreeNode{&TreeNode{Val: left}}
	}
	ret := []*TreeNode{}
	for i := left; i <= right; i++ {
		leftTrees := generateTrees(left, i-1)
		rightTrees := generateTrees(i+1, right)
		for j := range leftTrees {
			for k := range rightTrees {
				root := &TreeNode{Val: i}
				root.LChild = leftTrees[j]
				root.RChild = rightTrees[k]
				ret = append(ret, root)
			}
		}
	}
	return ret
}
