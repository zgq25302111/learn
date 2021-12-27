package main

import "fmt"

type TT struct {
	n int
}
func main() {
	m := make(map[int]TT)
	t := TT{7}
	m[0] = t
	fmt.Println(m[0].n)
}