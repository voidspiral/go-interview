package main

import (
	"fmt"
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
func main() {
	go worker()

	func() {
		for {
		}
	}()
}
