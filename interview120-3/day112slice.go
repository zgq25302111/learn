package main

import "fmt"

func main() {
	x := []int{100, 200, 300, 400, 500, 600, 700}	
	twohundred0 := &x[0]
	twohundred := &x[1]
	twohundred2 := &x[2]
	twohundred3 := &x[3]
	
	fmt.Println(twohundred0)
	fmt.Println(twohundred)
	fmt.Println(twohundred2)
	fmt.Println(twohundred3)
	
	x = append(x, 800)
	for i := range x {
		x[i]++
	}
	fmt.Println(x)
	fmt.Println(*twohundred)
}
