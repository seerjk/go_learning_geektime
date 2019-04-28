package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认零值 ""

	s = "hello"
	t.Log(len(s))

	//s[1] = '3'  // string是不可变的byte slice   cannot assign to s[1]
	//s = "\xE4\xB8\xA5" // stirng 可以存储任何二进制数据
	s = "\xE4\xB9\xAA\xFF"
	t.Log(s)

	/*
	* Unicode 一种字符集 code point
	* UTF8 是Unicode的一种存储实现（转换为字节序列列的规则）
	 */
	s = "中"
	t.Log(len(s)) // 是byte数

	// type rune = int32
	// It is used, by convention, to distinguish character values from integer values.
	c := []rune(s)
	t.Log(len(c))
	t.Logf("%x, %b", c, c)
	t.Log(c)                     // 十进制
	t.Logf("中 unicode %x", c[0]) // 十六进制
	t.Logf("中 UTF8 %x", s)       // 十六进制
}
