package main

import "fmt"

func Triple(n int) (r int) {
	defer func() {
		r += n // modify the return value
	}()

	return n + n // <=> r = n + n; return
}

func main() {
	fmt.Println(Triple(5)) // 15
}