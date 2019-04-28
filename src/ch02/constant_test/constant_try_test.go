package constant_test

import "testing"

// 连续常量定义简化方式
const (
	Monday    = iota + 1 // 1
	Tuesday              // 2
	Wednesday            // 3
	Thursday             // 4
)

const (
	Readable   = 1 << iota // 0001
	Writable               // 0010
	Executable             // 0100
)

func TestConstantTest(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 // 0111
	a = 1  // 0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	// a&Readable == 0111 & 0001 == 0001 == Readable
}
