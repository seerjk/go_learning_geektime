package operator_test

import "testing"

const (
	Readable   = 1 << iota // 0001
	Writable               // 0010
	Executable             // 0100
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// 长度不同数据，比较会编译错误
	// invalid operation: a == c (mismatched types [4]int and [5]int)
	//t.Log(a == c)

	// 元素 数量、值、顺序 都相同
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	/* &^ 按位置零操作符
	  1 &^ 0 -- 1
	  1 &^ 1 -- 0
	  0 &^ 1 -- 0
	  0 &^ 0 -- 0
	右操作数是1，则结果是0；右操作数是0，则结果是左操作数。
	 */
	a := 7 // 0111
	//a = 1  // 0001
	a = a &^ Readable // 清可读位
	a = a &^ Writable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	// a&Readable == 0111 & 0001 == 0001 == Readable
}
