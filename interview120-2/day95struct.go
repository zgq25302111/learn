package main

import "fmt"

type fooa struct{ Val int }
type bara struct{ Val int }
func main() {
	a := &fooa{Val: 5}
	b := &fooa{Val: 5}
	c := fooa{Val: 5}
	d := bara{Val: 5}
	e := bara{Val: 5}
	f := bara{Val: 5}
	fmt.Print(a == b, c == fooa(d), e == f)
}
