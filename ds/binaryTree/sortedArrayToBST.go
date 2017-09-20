package binaryTree

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToBST(nums, 0, len(nums)-1)
}

func sortedArrayToBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	if left == right {
		return &TreeNode{Val: nums[left]}
	}
	middle := (left + right) / 2
	root := &TreeNode{Val: nums[middle]}
	root.LChild = sortedArrayToBST(nums, left, middle-1)
	root.RChild = sortedArrayToBST(nums, middle+1, right)
	return root
}
