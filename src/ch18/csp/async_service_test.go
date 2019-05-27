package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	// 顺序执行
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// 异步返回

	// 无buffer 的 channel ，调用者和被调用者 都会阻塞在chan处，除非写入并读取chan的数据
	//retCh := make(chan string)

	// 有buffer的 channel，buffer满之前不会阻塞，会继续执行后续语句
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned result.")
		// write to chan retCh
		// 非buffer chan的阻塞处
		retCh <- ret
		fmt.Println("service exited.")
	}()

	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	// read from chan retCh
	// 非buffer chan的阻塞处
	fmt.Println(<-retCh)

	// wait until done
	time.Sleep(time.Second * 1)
}
