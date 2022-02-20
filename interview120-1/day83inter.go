package main

import "fmt"

type datac struct {
	name string
}
func (p *datac) print() {
	fmt.Println("name:", p.name)
}
type printer interface {
	print()
}
func main() {
	d1 := datac{"one"}
	d1.print()
}
