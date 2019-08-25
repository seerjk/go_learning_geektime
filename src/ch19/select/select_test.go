package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	//time.Sleep(time.Millisecond * 500)
	return "Done (service)"
}

func AsyncService() chan string {
	// 异步返回

	// 无buffer 的 channel ，调用者和被调用者 都会阻塞在chan处，除非写入并读取chan的数据
	//retCh := make(chan string)

	// 有buffer的 channel，buffer满之前不会阻塞，会继续执行后续语句
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned result. (AsyncService)")
		// write to chan retCh
		// 非buffer chan的阻塞处
		retCh <- ret
		fmt.Println("service exited. (AsyncService)")
	}()

	return retCh
}

func TestSelect(t *testing.T) {
	// 超时机制
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("timeout")
	}
}
