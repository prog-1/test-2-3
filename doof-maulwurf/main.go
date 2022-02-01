package main

import "fmt"

func doofmaulwurf(num []uint{1,0,0,0,0},row int) {
	fmt.Println(num)
	firstindx := 1
	secondindx := 0
	thirdindx := 0
	fourindx := 0
	fiveindx := 0
	if _,row > 0,row--{
		for i := range num{
			secondindx = num[1]
			thirdindx = num[2]
			fourindx = num[3]
			fiveindx = num[4]
			num[] += uint(firstindx,secondindx + firstindx,thirdindx + secondindx,fourindx + thirdindx,firstindx + fourindx)
			fmt.Println(num[])
		}
	}
}

func main() {
	var row uint
	fmt.Println("Enter rows: ")
	fmt.Scan(&row)
	fmt.Println(doofmaulwurf())
}
