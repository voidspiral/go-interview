package main

import "fmt"

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
