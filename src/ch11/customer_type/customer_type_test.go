// 自定义类型定义
package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

//func timeSpent(inner func(op int) int) func(op int) int {
//	// 计算函数操作的时长
//	return func(n int) int {
//		start := time.Now()
//		ret := inner(n)
//		fmt.Println("time spent:", time.Since(start).Seconds())
//		return ret
//	}
//}

func timeSpent(inner IntConv) IntConv {
	// 计算函数操作的时长
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
