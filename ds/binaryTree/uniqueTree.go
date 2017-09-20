package binaryTree

//	BST tree
// G(n) = F(1,n)+F(2,n)+...+F(n,n)
// F(i,n) = G(i-1) + G(n-i)
// http://www.cnblogs.com/pinganzi/p/6709867.html
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n+1; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[i-1] + G[n-i]
		}
	}
	return G[n]
}
