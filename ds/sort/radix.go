package sort

func RadixSort(n []int) {
	length := len(n)
	if length < 2 {
		return
	}
	radix := 1
	for {
		bucket := make([]int, 10)
		radixArray := []int{}
		flag := 0
		for _, i := range n {
			tmp := i % (10 * radix) / radix
			radixArray = append(radixArray, tmp)
			bucket[tmp]++
			if tmp != 0 {
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
		radix *= 10
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		preN := make([]int, length)
		copy(preN, n)
		for i := length - 1; i >= 0; i-- {
			bucket[radixArray[i]] -= 1
			n[bucket[radixArray[i]]] = preN[i]
		}
	}
}
