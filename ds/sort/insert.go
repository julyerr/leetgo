package sort

func Insert(n []int) {
	if n == nil {
		return
	}
	for i := 1; i < len(n); i++ {
		j := i
		for j > 0 {
			if n[j] < n[j-1] {
				n[j], n[j-1] = n[j-1], n[j]
			}
			j -= 1
		}
	}
}

func Insert1(n []int) {
	if n == nil {
		return
	}
	i, j := 0, 0
	for i = 1; i < len(n); i++ {
		if n[i] >= n[i-1] {
			continue
		}
		for j = 0; j < i; j++ {
			if n[i] <= n[j] {
				tmp := n[i]
				for k := i; k >= j+1; k-- {
					n[k] = n[k-1]
				}
				n[j] = tmp
				break
			}
		}
	}
}
