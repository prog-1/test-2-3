package main

import "fmt"

func filterdigits(){
	for i:=0, i<len(num),i++{
		if i = 1 || 3{
			if i+1 = 2, 4, 6, 8{
				num = len(num)-num[i]int
			}
		}
	}
	return
}

func main(){
	num := []int
	fmt.Println("Enter numbers: ")
	fmt.Scan(&num)
	return num
	fmt.Println(filterdigits()) 
}