package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		panic(1)
	}()
	//defer recover()
	panic(2)
}