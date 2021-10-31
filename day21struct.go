package main

import "fmt"

type SS struct {
	m string
}
func ff() *SS {
	return &SS{"foo"}
}
func main() {
	p := ff()
	fmt.Println(p.m) //print "foo"
}
