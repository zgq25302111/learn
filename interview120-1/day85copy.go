package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}
	dst = append(dst, src...)
	fmt.Println(dst)
}