package main

import "fmt"

type NN int
func (n NN) test(){
	fmt.Println(n)
}
func main() {
	var n NN = 10
	p := &n
	n++
	f1 := n.test
	n++
	f2 := p.test
	n++
	fmt.Println(n)
	f1()
	f2()
}
