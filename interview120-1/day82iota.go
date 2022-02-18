package main

import "fmt"

const (
	azero = iota
	aone = iota
)
const (
	infoa = "msg"
	bzero = iota
	bone = iota
)
func main() {
	fmt.Println(azero, aone)
	fmt.Println(bzero, bone)
}