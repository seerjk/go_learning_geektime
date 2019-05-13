package slice_test

import "testing"

// 切片内部结构  slice
/*
ptr *Elem
len int  // 元素的个数
cap int  // 内部数组的容量
*/

func TestSliceInit(t *testing.T) {
	var s0 []int // 不写长度的是slice，数组有长度
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	/* make([]type, len, cap)
	 len个元素会被初始化为默认零值
	未初始化元素不可以访问 */

	t.Log(len(s2), cap(s2))
	// panic: runtime error: index out of range [recovered]
	// t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	// len 可访问，初始化的元素的个数
	// cap 总容量，不一定都能访问
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSlice(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
		// 当cap不足，len==cap后继续append元素，cap会增长为2倍当原有的cap
		// 为什么不写作 append(s, i) 因为slice存储空间不足后，或申请一段为原来2倍大的新的连续空间，空间起始地址变了，需要重新赋值。
		// 写作 s = append(s, i)
	}
}

func TestSliceShareMemory(t *testing.T) {
	// slice会共享内存

	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

/*
数组vs切片

1。 容量是否可伸缩
	数组不可
	slice可以

2。 是否可以进行比较
	array 相同维数，相同元素 可以比较
	slice 不可以比较
*/
func TestSliceCompare(t *testing.T) {
	// slice
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}

	// a == b (slice can only be compared to nil)
	//if a == b {
	//	t.Log("a==b")
	//}

	// array 数组  a := [...]int{1, 2, 3, 4}
	// slice 切片  a := []int{1, 2, 3, 4}
}
