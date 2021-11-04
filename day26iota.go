package main

import "fmt"

const (
	aa = iota
	b = iota
)
const (
	name = "name"
	c = iota
	d = iota
)
func main() {
	fmt.Println(aa)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}