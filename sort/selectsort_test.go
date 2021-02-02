package sort

import (
	"log"
	"testing"
)

func TestSelectSort(t *testing.T) {
	type args struct {
		arr []int
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.arr)
			log.Println(tt.args.arr)
		})
	}
}
