package main

import "fmt"

func main() {
	x := []int{
		1,
		2,
	}
	x = x
	fmt.Println(x)

	y := []int{3,4,}
	_ = y
}
