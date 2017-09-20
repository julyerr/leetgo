package sort

func Bubble(n []int) {
	if n == nil {
		return
	}
	for i := 0; i < len(n)-1; i++ {
		tmp := n[i]
		for j := i + 1; j < len(n); j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
		if tmp == n[i] {
			break
		}
	}
}
