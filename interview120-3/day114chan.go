package main

import "fmt"

func main() {
	i := 0
	f := func() int {
		fmt.Println("incr")
		i++
		return i
	}
	c := make(chan int)
	for j := 0; j < 2; j++ {
		select {
		case c <- f():
			
		default:
			fmt.Println("incr+")
		}
	}
	fmt.Println(i)
}
