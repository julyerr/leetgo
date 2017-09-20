package sort

import (
	st "sort"
)

type Interval struct {
	Start int
	End   int
}

type TL []Interval

func (t TL) Len() int {
	return len(t)
}

func (t TL) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TL) Less(i, j int) bool {
	if t[i].Start == t[j].Start {
		return t[i].End < t[j].End
	}
	return t[i].Start < t[j].Start
}

func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	// 设置TL，赋值的时候强制转换
	var il TL
	il = intervals
	st.Sort(il)
	ret := []Interval{}
	for i := 0; i < len(il); {
		// [2,4] [4,5]
		for i != len(il)-1 && il[i].End >= il[i+1].Start {
			il[i+1].Start = il[i].Start
			if il[i+1].End < il[i].End {
				// [2,5] [3,4]
				il[i+1].End = il[i].End
			}
			i++
		}
		ret = append(ret, il[i])
		i++
	}
	return ret
}
