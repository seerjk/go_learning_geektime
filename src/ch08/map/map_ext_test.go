package map_ext_test

import "testing"

// 工厂模式
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

// 实现set  map[type]bool
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true

	n := 3
	if mySet[n] {
		t.Logf("%d is existing.", n)
	} else {
		t.Logf("%d is not existing.", n)
	}

	// 输出集合中独立元素的个数
	mySet[3] = true
	t.Log(len(mySet))

	// 从set中删除元素
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing.", n)
	} else {
		t.Logf("%d is not existing.", n)
	}
}

func ExistInSet(n int, mySet map[int]bool) bool {
	if mySet[n] {
		return true
	} else {
		return false
	}
}
