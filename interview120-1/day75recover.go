package main

import "fmt"

func fb() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}
func main() {
	fb()
}
