package sort

func SortColors(n []int) {
	length := len(n)
	if length < 2 {
		return
	}
	colorArrays := make([]int, 3)
	for _, v := range n {
		colorArrays[v]++
	}
	j := 0
	for i := 0; i <= 2; i++ {
		for colorArrays[i] > 0 {
			n[j] = i
			j += 1
			colorArrays[i]--
		}
	}
}
