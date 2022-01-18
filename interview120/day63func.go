package main

import "fmt"

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}
func main() {
	f := F(5)
	defer func() {
		fmt.Println(f())
	}()
	defer fmt.Println(f())
	i := f()
	fmt.Println(i)
}