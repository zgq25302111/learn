package main

import "fmt"

const (
	Century = 100
	Decade = 010
	Year = 001
)
func main() {
	fmt.Println(Century + 2*Decade + 2*Year)
}