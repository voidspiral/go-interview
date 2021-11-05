package fanin

import (
	"fmt"
	"time"
)

func main() {
	receive()
}
func receive() {
	ch := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
	for {
		i, ok := <-ch
		fmt.Println(i, ok)
		time.Sleep(time.Second)
	}
}
