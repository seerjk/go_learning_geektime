package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	// 协程传递变量的正确示例：
	// go的函数传递的值，进行了值的拷贝，主协程和子协程中的i不存在共享
	// i传递到子协程时，进行了值拷贝
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	// 并发编程容易错误的示例：
	// 输出结果全部都是 10，与预期结果不一致
	// go func(){}() 协程中直接调用了主协程的变量i，当作共享变量使用
	// 存在竞争关系，需要用 lock

	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}

	// 等待执行完毕
	time.Sleep(time.Millisecond * 50)
}
