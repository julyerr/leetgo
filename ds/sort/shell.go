package sort

func ShellSort(n []int) {
	length := len(n)
	for step := length / 2; step > 0; step /= 2 {
		for i := length - 1; i-step >= 0; i-- {
			for j := i; j-step >= 0; j -= step {
				tmp := j - step
				if n[j] < n[tmp] {
					n[j], n[tmp] = n[tmp], n[j]
				}
			}
		}
	}
}
