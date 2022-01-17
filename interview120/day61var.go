package main

import "fmt"

func main() {
	var k = 1
	var s = []int{1, 2}
	k, s[k] = 0, 3
	fmt.Println(k)
	fmt.Println(s[1])
}