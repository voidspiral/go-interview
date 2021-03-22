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
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			//执行定时任务
			fmt.Println("执行1s定时任务")
		}
	}
}
func main() {

}
