package sort

func MergeSort(n []int) []int {
	length := len(n)
	if length <= 1 {
		return n
	}
	middle := length / 2
	left := MergeSort(n[:middle])
	right := MergeSort(n[middle:])
	return mergeSortedArrays(left, right)
}

func mergeSortedArrays(m, n []int) []int {
	i, j := 0, 0
	ret := []int{}
	for i < len(m) && j < len(n) {
		if m[i] < n[j] {
			ret = append(ret, m[i])
			i++
		} else {
			ret = append(ret, n[j])
			j++
		}
	}
	ret = append(ret, m[i:]...)
	ret = append(ret, n[j:]...)
	return ret
}
