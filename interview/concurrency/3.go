package concurrency

import "sync"

// 交替打印字符串
var s = "hello world"

func solve3() {
	done := make(chan struct{})
	done2 := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < len(s); i += 2 {
			<-done
			println(string(s[i]))
			if i >= len(s) {
				return
			}
			done2 <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < len(s); i += 2 {
			<-done2
			println(string(s[i]))

			done <- struct{}{}
		}
	}()
	done <- struct{}{}
	wg.Wait()
}
