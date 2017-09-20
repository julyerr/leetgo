package sort

func HIndex(citation []int) int {
	length := len(citation)
	if length == 0 {
		return 0
	}
	array := make([]int, length+1)
	for i := 0; i < length; i++ {
		if citation[i] > length {
			array[length] += 1
		} else {
			array[citation[i]] += 1
		}
	}
	t := 0
	for i := length; i > 0; i-- {
		t += array[i]
		if t >= i {
			return i
		}
	}
	return 0
}
