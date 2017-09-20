//介绍和实现： https://songlee24.github.io/2015/05/09/prefix-tree/
//starts most: https://github.com/derekparker/trie
package trie

import (
	"fmt"
)

// to show chinese well,use rune rather string

type node struct {
	child map[rune]*node
	isEnd bool
}

func createNode() *node {
	return &node{
		child: make(map[rune]*node),
		isEnd: false,
	}
}

var (
	root *node
)

func Insert(txt string) {
	length := len(txt)
	if length < 1 {
		return
	}
	if root == nil {
		root = createNode()
	}
	keys := []rune(txt)
	n := root
	for _, v := range keys {
		if _, ok := n.child[v]; !ok {
			n.child[v] = createNode()
		}
		n = n.child[v]
	}
	n.isEnd = true
}

func Select() {
	dfs(root, []rune{})
}

func dfs(root *node, str []rune) {
	if root == nil {
		return
	}
	for k, v := range root.child {
		s := append(str, k)
		if v.isEnd {
			fmt.Println(string(s))
		}
		dfs(v, s)
	}
}
