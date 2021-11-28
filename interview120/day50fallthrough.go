package main

import "fmt"

func main() {
	handle(0)
}

func handle(i int) {
	switch i {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}
