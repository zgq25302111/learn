package main

import "fmt"

type Ta struct{}

func (*Ta) foo() {
}
func (Ta) bar() {
}
type Sa struct {
	*Ta
}
func main() {
	s := Sa{}
	_ = s.foo
	s.foo()
	//_ = s.bar
	fmt.Printf("%#v",s)
}