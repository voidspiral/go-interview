package main

import (
	"fmt"
	"time"
)

func main() {
	receicer()
}
func receicer() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	for {
		i, ok := <-ch
		fmt.Println(i, ok)
		time.Sleep(time.Second)
	}
}
