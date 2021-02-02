package sort

import (
	"log"
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	test4 := args{}
	//rand.Seed(42)
	for i := 0; i < 5; i ++ {
		test4.arr = append(test4.arr, rand.Intn(1000))
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test case 1:",
			args: args{[]int{1, 2, -100, 4, 5}},
		},
		{
			name: "test case 2:",
			args: args{[]int{7, 6, 5, 2, 100}},
		},
		{
			name: "test case 3:",
			args: args{[]int{-100,-100,0,0,1,-1}},
		},
		//{
		//	name:"test case 4",
		//	args: test4,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			BubbleSort(tt.args.arr)
			log.Println(tt.name, tt.args.arr)
		})
	}
}
