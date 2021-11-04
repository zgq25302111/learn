//当且仅当动态值和动态类型都为 nil 时

package main

import "fmt"

type People interface {
	Show()
}
type Student struct{}
func (stu *Student) Show() {
}
func main() {
	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
