package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
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
	// _  忽略其中一个返回值
	a, _ := returnMultiValues()
	t.Log(a)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 可变长参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

// defer 延迟执行函数
// 清理资源，释放锁等。panic后也会执行defer
func Clear() {
	fmt.Println("Clear resources.")
}
func TestDefer(t *testing.T)  {
	defer Clear()
	fmt.Println("Start")
	panic("err") // defer 也会执行, 安全的释放资源和锁

	// panic 后面的代码不可达
	fmt.Println("End")
}
