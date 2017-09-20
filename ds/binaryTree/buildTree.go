package binaryTree

import "fmt"

// if for small dataset
func BuildTreePI1(pre []int, in []int) *TreeNode {
	length := len(pre)
	if length == 0 {
		return nil
	}
	pos := 0
	for in[pos] != pre[0] {
		pos++
	}
	root := &TreeNode{Val: pre[0]}
	root.LChild = BuildTreePI1(pre[1:pos+1], in[:pos])
	root.RChild = BuildTreePI1(pre[pos+1:], in[pos+1:])
	return root
}

// 递归代码嵌套层次少，但是条件判断较多
// 代码简洁，嵌套层次少，综合来看，还是简洁更为合适
func BuildTreeAI(post []int, in []int) *TreeNode {
	length := len(post)
	if length == 0 {
		return nil
	}
	pos := 0
	for in[pos] != post[length-1] {
		pos++
	}
	root := &TreeNode{Val: post[length-1]}
	if pos == 0 {
		root.LChild = nil
	} else {
		root.LChild = BuildTreeAI(post[:pos], in[:pos])
	}
	if pos == length-1 {
		root.RChild = nil
	} else {
		root.RChild = BuildTreeAI(post[pos:length-1], in[pos+1:])
	}
	return root
}

func BuildTreePI(pre []int, in []int) *TreeNode {
	// no modre duplicated arrays, try to using map struct to improve efficiency
	m := make(map[int]int)
	for i := range in {
		m[in[i]] = i
	}
	return BuildTreePITmp(pre, in, m, 0)
}

func BuildTreePITmp(pre []int, in []int, m map[int]int, base int) *TreeNode {
	length := len(pre)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return &TreeNode{Val: pre[0]}
	}
	pos := m[pre[0]] - base
	fmt.Println(pos)
	root := &TreeNode{Val: pre[0]}
	root.LChild = BuildTreePITmp(pre[1:pos+1], in[:pos], m, base)
	root.RChild = BuildTreePITmp(pre[pos+1:], in[pos+1:], m, pos)
	return root
}
