package binaryTree

var max int

func FindFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	m := make(map[int]int)
	ret := []int{}
	// 记得给全局变量赋初始值
	max = 0
	findFrequentTreeSum(root, m)
	for k, v := range m {
		if v == max {
			ret = append(ret, k)
		}
	}
	return ret
}

func findFrequentTreeSum(root *TreeNode, m map[int]int) int {
	if root == nil {
		return 0
	}
	s1 := findFrequentTreeSum(root.LChild, m)
	s2 := findFrequentTreeSum(root.RChild, m)
	sum := s1 + s2 + root.Val
	m[sum] += 1
	if m[sum] > max {
		max = m[sum]
	}
	return sum
}
