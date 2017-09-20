// package trie

// import (
// 	"fmt"
// )

// type node struct {
// 	child map[rune]*node
// 	isEnd bool
// }

// var (
// 	root *node
// )

// func newNode() *node {
// 	return &node{
// 		child: make(map[rune]*node),
// 		isEnd: false,
// 	}
// }

// func Insert(txt string) {
// 	if len(txt) < 1 {
// 		return
// 	}

// 	key := []rune(txt)
// 	if root == nil {
// 		root = newNode()
// 	}
// 	n := root
// 	for i := 0; i < len(key); i++ {
// 		if _, ok := n.child[key[i]]; !ok {
// 			n.child[key[i]] = newNode()
// 		}
// 		n = n.child[key[i]]
// 	}
// 	n.isEnd = true
// }

// func Select() {
// 	dfs(root, []rune{})
// }

// func dfs(n *node, str []rune) {
// 	if n == nil {
// 		return
// 	}
// 	for k, v := range n.child {
// 		s := append(str, k)
// 		if v.isEnd == true {
// 			fmt.Println(string(s))
// 		}
// 		dfs(v, s)
// 	}
// }
