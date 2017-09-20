package binaryTree

import "fmt"

func LevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		array := []*TreeNode{}
		values := []int{}
		for i := 0; i < len(queue); i++ {
			values = append(values, queue[i].Val)
			if queue[i].LChild != nil {
				array = append(array, queue[i].LChild)
			}
			if queue[i].RChild != nil {
				array = append(array, queue[i].RChild)
			}
		}
		fmt.Println(values)
		queue = array
		ret = append(ret, values)
	}
	return ret
}

func ZigZagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	ret := [][]int{}
	flag := 0
	for len(queue) > 0 {
		arrays := []*TreeNode{}
		Values := []int{}
		for i := 0; i < len(queue); i++ {
			Values = append(Values, queue[i].Val)
			if queue[i].LChild != nil {
				arrays = append(arrays, queue[i].LChild)
			}
			if queue[i].RChild != nil {
				arrays = append(arrays, queue[i].RChild)
			}
		}
		if flag%2 != 0 {
			length := len(Values)
			for i := 0; i < length/2; i++ {
				Values[i], Values[length-1-i] = Values[length-1-i], Values[i]
			}
		}
		queue = arrays
		ret = append(ret, Values)
		flag = flag + 1
	}
	return ret
}
