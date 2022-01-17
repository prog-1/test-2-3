package main

import (
	"fmt"
	"testing"
)

func Testdoofmaulwurf(t *testing.T) {
	for _, tc := range []struct {
		name    string
		row		uint
		num   []int
		want   []int
	}{
		{"empty",0 ,[]int{},[]int{}},
		{"1",1 ,[]int{1 0 0 0 0},[]int{1 0 0 0 0}},


	} {
		if got := doofmaulwurf(tc.num,tc.row); got != tc.want {
			fmt.Printf("FAIL %s: got %v, want %v", tc.num,tc.row got, tc.want)

		}
	}
}