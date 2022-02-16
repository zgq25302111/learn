package main

import "fmt"

type TimesMatcher struct {
	base int
}
func NewTimesMatcher(base int) *TimesMatcher {
	return &TimesMatcher{base:base}
}
func main() {
	p := NewTimesMatcher(3)
	fmt.Println(p)
}