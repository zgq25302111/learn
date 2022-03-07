package main

import "fmt"

var xa int
func init() {
	xa++
}
func main() {
	fmt.Println(xa)
}
