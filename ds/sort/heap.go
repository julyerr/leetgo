package sort

func HeapSort(n []int) {
	length := len(n)
	middle := length / 2
	// find the middle, whether it has one child or two
	for i := middle; i >= 0; i-- {
		heap(n, i, length-1)
	}
	// at least two child
	for i := length - 1; i > 0; i-- {
		n[0], n[i] = n[i], n[0]
		heap(n, 0, i-1)
	}
}

func heap(n []int, start, end int) {
	left := start*2 + 1
	if start > left {
		return
	}
	right := left + 1
	max := left
	if right <= end && n[right] > n[left] {
		max = right
	}
	if n[start] >= n[max] {
		return
	}
	n[start], n[max] = n[max], n[start]
	heap(n, max, end)
}
