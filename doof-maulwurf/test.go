package main

import (
	"fmt"
	"testing"
)

func Testdoofmaulwurf(t *testing.T) {
	for _, tc := range []struct {
		name    string
		row		uint
		want   []int
	}{
		{"empty",0 ,[]int{},[]int{}},
		{"1",1 ,[]int{1 0 0 0 0}},
		{"3",3 ,[]int{1 0 0 0 0
			1 1 0 0 0
			1 2 1 0 0}},
		{"5",0 ,[]int{1 0 0 0 0
			1 1 0 0 0
			1 2 1 0 0
			1 3 3 1 0
			1 4 6 4 1}},
		{"negative",-1 ,[]int{}},
		{"8",0 ,[]int{1 0 0 0 0
			1 1 0 0 0
			1 2 1 0 0
			1 3 3 1 0
			1 4 6 4 1
			1 5 10 10 1
			1 6 15 20 1
			1 7 21 35 1}},

	} {
		if got := doofmaulwurf(tc.row); got != tc.want {
			fmt.Printf("FAIL %s: row %v, want %v", tc.row got, tc.want)

		}
	}
	if ok {
		fmt.Println("PASS")
	}
}