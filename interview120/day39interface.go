package main

import "fmt"

func Foo2(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int
	Foo2(x)
}
