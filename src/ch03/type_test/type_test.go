package type_test

import (
	"math"
	"testing"
)

type MyInt int64


func TestImplicit(t *testing.T) {
	// 不支持隐式类型转换，即使别名也不支持
	// 别名到原类型

	var a int32 = 1
	var b int64

	// cannot use a (type int) as type int64 in assignment
	// cannot use a (type int32) as type int64 in assignment
	//b = a

	b = int64(a)

	var c MyInt
	// cannot use b (type int64) as type MyInt in assignment
	//c = b

	c = MyInt(b)
	t.Log(a, b, c)
}

func TestMaxValue(t *testing.T) {
	// 类型的预定义值
	t.Log(math.MaxInt64)
	t.Log(math.MaxFloat64)
	t.Log(math.MaxUint32)
}

func TestPoint(t *testing.T) {
	// 指针类型

	a := 1
	aPtr := &a
	// 1. 不支持指针运算
	// invalid operation: aPtr + 1 (mismatched types *int and int)
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	// string 是值类型 默认初始化为 "" 空字符串 而不是 nil
	t.Log("*"+s+"*")
	t.Log(len(s))
	if s == "" {
		t.Log("s is empty.")
	}
}