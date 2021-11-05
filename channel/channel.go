package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second)
		ch <- 100
	}()
	//fmt.Println("channle:%v\n, <-ch")

	select {
	case v, _ := <-ch:
		fmt.Println(v)
	}

}
