package main

import (
	"fmt"
	"unsafe"
)

//type slice struct {
//	ptr uintptr
//	len int
//	cap int
//}

func main() {
	someSlice := make([]byte, 666)

	fmt.Println((unsafe.Sizeof(&someSlice)))

}