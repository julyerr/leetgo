// http://blog.csdn.net/wincol/article/details/4795369
// http://blog.csdn.net/q547550831/article/details/51860017

// http://blog.cyeam.com/golang/2014/08/08/go_index

package strings

// 将所有需要比较的字符预先存入map结构，可以考虑unicode以便支持中文搜索
// 字符数越多，出现赋值的情况越多，效率低小
func Sunday(origin, pattern string) int {
	m := make(map[byte]int)
	lLen := len(origin)
	rLen := len(pattern)
	i := 0
	for i < 256 {
		m[byte(i)] = rLen + 1
		i += 1
	}
	for i := 0; i < rLen; i++ {
		m[pattern[i]] = rLen - i
	}
	start := 0
	for start <= lLen-rLen {
		j := 0
		for origin[start+j] == pattern[j] {
			j += 1
			// rLen >= 1
			if j >= rLen {
				return j
			}
		}
		start += m[origin[start+rLen]]
	}
	return -1
}
