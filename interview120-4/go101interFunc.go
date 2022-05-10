package main

import "fmt"

type I interface {
	m(int)bool
}

type T string
func (t T) m(n int) bool {
	return len(t) > n
}

func main() {
	var i I = T("go")
	fmt.Println(i.m(5))                        // true
	fmt.Println(I.m(i, 5))                     // true
	fmt.Println(interface{m(int)bool}.m(i, 5)) // true
}