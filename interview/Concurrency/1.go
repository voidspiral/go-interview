package Concurrency

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// 题目： n个goroutine 交替打印0~n

func solve() {
	const n int = 4
	const maxNum int = 100
	ch := make([]chan struct{}, n)
	for i := range ch {
		ch[i] = make(chan struct{})
	}
	wg := sync.WaitGroup{}
	wg.Add(n)
	i := 0
	go func() {
		ch[0] <- struct{}{} //一个开始的信号
	}()
	for j := 0; j < n; j++ {
		chanNum := (j + 1) % n
		go func() {
			defer wg.Done()
			for {
				<-ch[chanNum]
				fmt.Printf("i am goroutine %d\n", chanNum+1)
				fmt.Println(i) //顺带练习n个goroutine交替打印 0 ~ n
				if i >= math.MaxInt64-1 {
					return
				}
				i++
				time.Sleep(time.Second) //控制打印速率
				ch[(chanNum+1)%n] <- struct{}{}

			}
		}()
	}
	wg.Wait()
}
