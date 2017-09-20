package sort

// 对于训练和学习需要合适的度量
// 先给自己适当的思考和动手实践
// 过了时间，然后参考和学习
// 大体理解之后，按照自己的理解重新手写
// 调整发现问题，修改
// 重新书写总结心得
func QuickSort(n []int) {
	length := len(n)
	if length <= 1 {
		return
	}
	head, tail := 0, length-1
	middle := n[0]
	for i := 1; i <= tail; {
		if n[i] > middle {
			n[i], n[tail] = n[tail], n[i]
			tail--
		} else {
			n[i], n[head] = n[head], n[i]
			head++
			i++
		}
	}
	QuickSort(n[:head])
	QuickSort(n[head+1:])
}
