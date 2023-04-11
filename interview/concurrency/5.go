package concurrency

import (
	"fmt"
	"runtime"
)

//打印什么结果
//考察两点，但我不知道做什么
func solve5() {
	//runtime.GOMAXPROCS(1)
	var i byte
	go func() {
		for i = 0; i < 255; i++ {
			fmt.Printf("index_i：%v\n", i)
		}
	}()
	fmt.Println("嘿")
	runtime.GC()
	fmt.Println("哈")
}
