package main

import "fmt"

type Nn int
func (n *Nn) test(){
	fmt.Println(*n)
}
func main() {
	var n Nn = 10
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