package sort

func Select(n []int) {
	if n == nil {
		return
	}
	for i := 0; i < len(n)-1; i++ {
		pos, min := i, n[i]
		for j := i + 1; j < len(n); j++ {
			if n[j] < min {
				pos, min = j, n[j]
			}
		}
		n[i], n[pos] = n[pos], n[i]
	}
}
