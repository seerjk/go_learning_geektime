package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounterUnsafe(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterSafe(t *testing.T) {
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

func TestCounterSafeWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		// 每启动一个 groutine前 wg 加一
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			// 完成一个 groutine后 done
			wg.Done()
		}()
	}

	//time.Sleep(1 * time.Second)
	// 等待全部任务完成，阻塞在这里，等所有任务都完成
	wg.Wait()
	t.Logf("counter = %d", counter)
}

// rwlock 和 mutex 区别
// rwlock 读写锁分开 效率更高
// mutex 互斥锁 效率低
