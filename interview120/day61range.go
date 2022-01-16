package main

import "fmt"

func main() {
	var k = 9
	for k = range []int{} {
		fmt.Println(k)
	}
}