package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[1], arr[2])

	arr1 := [4]int{1, 2, 3, 4}
	// 不想写长度 初始化时下表用 ...
	arr3 := [...]int{1, 2, 3, 5}
	arr1[1] = 5
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, v := range arr3 {
		t.Log(idx, v)
	}

	for _, v := range arr3 {
		t.Log(v)
	}
}

func TestMultidimensionalArray(t *testing.T) {
	// 多维数组
	arr4 := [2][2]int{{1, 2}, {3, 4}}
	for _, a := range arr4 {
		for _, b := range a {
			t.Log(b)
		}
	}
}

// 数组截取
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)

	arr3_sec1 := arr3[3:]
	// arr3[-1:]  不支持负数下标
	t.Log(arr3_sec1)

	t.Log(arr3[:])
}

