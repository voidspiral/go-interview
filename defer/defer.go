package main

import "fmt"

type number int
func (n *number) pprint(){
	fmt.Println(n)
}
func (n number) print() {
	fmt.Println(n)
}

func main() {
	//var whaterver [3] struct{}
	//for i := range whaterver {
	//	defer func(i int) {
	//		fmt.Println(i)
	//	}(i)
	//}
	var n number
	////defer n.print()
	////defer n.pprint()
	////defer func() {n.print()}()
	////defer func() {n.pprint()}()
	//defer fmt.Println("1", n)
	//defer func(){
	//	fmt.Println("2",n)
	//}()
	n = 3
	defer func() {
		n = n + 5
	}()
	fmt.Println(n)

	return
}

func f1() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func p3(){
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
}
//func accumulator() (err error){
//	defer func(){
//		if err := recover(); err != nil {
//
//		}
//	}()
//}