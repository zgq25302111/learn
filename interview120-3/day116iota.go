package main

import "fmt"

const (
 greeting = "Hello, Go"
 heloo = "2, Go"
 one = iota
 two
)
func main() {
 fmt.Println(one, two)
}