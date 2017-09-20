package binaryTree

// 直接考虑有多种情况的前提下，使用递归
// 递归有 bottom -> top 也有 top -> bottom
// bottom -> top:考虑返回值传递给调用程序
// top -> bottom 传入参数，一般不需要返回值，可以使用全局变量或者struct保留
type solution struct {
	sum int
}

func SumNumbers(root *TreeNode) int {
	this := &solution{}
	this.sumNumbers(root, 0)
	return this.sum
}

func (s *solution) sumNumbers(root *TreeNode, current int) {
	if root == nil {
		return
	}
	current = current*10 + root.Val
	if root.LChild == nil && root.RChild == nil {
		s.sum += current
	}
	s.sumNumbers(root.LChild, current)
	s.sumNumbers(root.RChild, current)
}

func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.LChild == nil && root.RChild == nil {
		return sum == 0
	}
	return HasPathSum(root.LChild, sum) || HasPathSum(root.RChild, sum)
}

func PathSum1(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	sum -= root.Val
	if root.LChild == nil && root.RChild == nil {
		if sum == 0 {
			return [][]int{{root.Val}}
		}
		return nil
	}
	s1 := PathSum1(root.LChild, sum)
	s2 := PathSum1(root.RChild, sum)
	for i := range s2 {
		s1 = append(s1, s2[i])
	}
	for i := range s1 {
		tmp := []int{root.Val}
		for j := range s1[i] {
			// append can only push to the right
			tmp = append(tmp, s1[i][j])
		}
		s1[i] = tmp
	}
	return s1
}

var ret [][]int

func PathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	onePath := []int{}
	pathSum(root, sum, onePath)
	return ret
}

func pathSum(root *TreeNode, sum int, onePath []int) {
	if root == nil {
		return
	}
	sum -= root.Val
	onePath = append(onePath, root.Val)
	if root.LChild == nil && root.RChild == nil {
		if sum == 0 {
			// add new array space
			tmp := make([]int, len(onePath))
			copy(tmp, onePath)
			ret = append(ret, tmp)
		}
	}
	pathSum(root.LChild, sum, onePath)
	pathSum(root.RChild, sum, onePath)
}
