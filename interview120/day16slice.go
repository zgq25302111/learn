package main

import "fmt"

type A interface {
	ShowA() int
}
type B interface {
	ShowB() int
}
type Work struct {
	i int
}
func (w Work) ShowA() int {
	return w.i + 10
}
func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	/*c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowB())
	fmt.Println(b.ShowA())

	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
	*/
	s := Work{3}
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
