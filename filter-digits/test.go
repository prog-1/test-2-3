package main

import (
	"fmt"
	"testing"
)

func Testfilterdigits(t *testing.T) {
	for _, tc := range []struct {
		name    string
		num   []int
		want   []int
	}{
		{"without 1 and 3", []int{69758},[]int{69758}},
		{"with 1 and 3", []int{1697538},[]int{69758}},
		{"with 1 and 3,but without positive numbers", []int{6193758},[]int{6193758}},
		{"empty", []int{},[]int{}},
		{"symbols", []int{++-+&*&^%$#},[]int{++-+&*&^%$#}},
		{"negative numbers", []int{-69758},[]int{-69758}},
		{"negative 1 and 3 with positive numbers", []int{-1697538},[]int{-69758}},

	} {
		if got := filterdigits(tc.num); got != tc.want {
			fmt.Printf("FAIL %s: got %v, want %v", tc.num, got, tc.want)

		}
	}
}