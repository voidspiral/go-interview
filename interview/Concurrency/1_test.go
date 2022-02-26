package Concurrency

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	//solve()
	ch := make(chan struct{})
	ch <- struct {
	}{}
	for i := 0; i < 10; i++ {
		go func() {
			<- ch
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
