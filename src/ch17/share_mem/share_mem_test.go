package share_mem

import (
	"sync"
	"testing"
	"time"
)

/*
并发控制：传统共享内存的方式
*/

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		// 并发竞争，不是线程安全的
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	// 加锁
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()

			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 等待其他协程完成 WaitGroup
func TestWaitGroup(t *testing.T) {
	// 加锁
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 5000; i++ {
		// 每增加一个协程，wg加一
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()

			counter++
			// 每一个协程完成后，告诉wg 完成了 Done
			wg.Done()
		}()
	}

	//time.Sleep(1 * time.Second)
	// wg 等待全部协程完成，阻塞在这里
	wg.Wait()

	t.Logf("counter = %d", counter)
}


// rw lock  readlock 非互斥  writelock 互斥锁