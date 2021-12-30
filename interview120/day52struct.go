package main

type X struct {}
func (x *X) test() {
	println(x)
}
func main() {
	var a *X
	a.test()
	x := X{}
	x.test()
}