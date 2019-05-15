package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	// cap=10

	t.Logf("len m3=%d", len(m3))
	// map 无cap
	// invalid argument m3 (type map[int]int) for cap
	//t.Log("cap m3:", cap(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	// 访问key不存在时，返回的值 0
	t.Log(m1[1])

	// 赋值为 0
	m1[2] = 0
	t.Log(m1[2])

	m1[3] = 0
	// map的key不存在时
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}

	// index, value := range array
	// key, value := range map
}