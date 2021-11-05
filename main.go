package main

import (
	"fmt"
	"sync"
	"time"
)

//func test() bool {
//	select {
//	case <- time.After(100 * time.Millisecond):
//	case <-s.stopc:
//		return false
//	}
//}

func worker() {
	ticker := time.NewTimer(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			//执行定时任务
			fmt.Println("执行1s定时任务")
			if !ticker.Stop() {

				<-ticker.C
			}
			ticker.Reset(time.Second)
		}
	}
}

const (
	a = 1 << iota
	b
	c = iota
)

func main() {
	//go worker()
	//
	//func() {
	//	for {
	//	}
	//}()
	fmt.Println(a, b, c)
	count := 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	m := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			// 对变量count进行10次加1
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
