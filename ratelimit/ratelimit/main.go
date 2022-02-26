package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	lock := sync.Mutex{}
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		count++
		go func() {

			fmt.Println(i)
		}()
	}

}
