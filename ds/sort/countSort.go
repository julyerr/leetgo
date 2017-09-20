package sort

func CountSort(n []int) []int {
	length := len(n)
	if length < 2 {
		return n
	}
	min, max := n[0], n[0]
	for i := 1; i < length; i++ {
		if min > n[i] {
			min = n[i]
		}
		if max < n[i] {
			max = n[i]
		}
	}
	newN := make([]int, max-min+1)
	for _, v := range n {
		newN[v-min]++
	}
	ret := []int{}
	for i, v := range newN {
		for v > 0 {
			ret = append(ret, i+min)
			v--
		}
	}
	return ret
}
