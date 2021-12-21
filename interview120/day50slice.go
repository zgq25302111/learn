package main

import "fmt"

type T struct {
	ls []int
}
func fooo(t T) {
	t.ls[0] = 100
}
func main() {
	var t = T{
		ls: []int{1, 2, 3},
	}
	fooo(t)
	fmt.Println(t.ls[0])
}
