package main

import "fmt"

func main() {
	isMatch := func(i int) bool {
		switch(i) {
		case 1:
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(2))
}