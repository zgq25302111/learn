package main

import "fmt"

func fn(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()
	var fn func()
	defer fn()
	fn = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(fn(3))
}