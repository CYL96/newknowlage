package atomic

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var key int64

func TestOne() {
	// 通过这个简单的测试得到了totalProcess是线程安全的
	var totalProcess int64 = 10000
	atomic.StoreInt64(&key, totalProcess)
	var i int64
	for i = 0; i < totalProcess; i++ {
		go func(n int64) {
			time.Sleep(time.Second + time.Duration(rand.Intn(10))*time.Second)
			atomic.AddInt64(&key, -1)
		}(i)
	}
	for {
		process := atomic.LoadInt64(&key)
		if process == 0 {
			break
		} else {
			fmt.Println(process, " is running")
		}

		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("all done")
}

func TestOneWithOutAtomic() {
	// 通过这个简单的测试得到了totalProcess并不是线程安全的
	var totalProcess int64 = 10000
	var i int64
	for i = 0; i < totalProcess; i++ {
		go func(n int64) {
			time.Sleep(time.Second + time.Duration(rand.Intn(10))*time.Second)
			totalProcess--
		}(i)
	}
	for {
		if totalProcess == 0 {
			break
		} else {
			fmt.Println(totalProcess, " is running")
		}

		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("all done")
}
