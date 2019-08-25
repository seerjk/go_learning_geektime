package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)

		// 向关闭的chan发消息 panic: send on closed channel
		//ch <- 1

		wg.Done()
	}()
}

// 消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// for i := 0; i < 12; i++ {
			// data := <-ch
			// fmt.Println(data)

			if data, ok := <-ch; ok {
				// true chan 正常接收状态
				fmt.Println(data)
			} else {
				// false 通道关闭, 退出
				//fmt.Println("chan closed.")
				break
			}
		}
		wg.Done()
	}()
}

// CSP 通过chan链接 生产者和消费者
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)
	//wg.Add(1)
	//dataReceiver(ch, &wg)
	wg.Wait()
}
