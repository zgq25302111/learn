package main

import "fmt"

func main() {
	var x uint8 = 214
	var y uint8 = 92
	fmt.Printf("x : %08b\n",x)
	fmt.Printf("y : %08b\n",y)
	//fmt.Printf("| : %08b\n",x | y)
	fmt.Printf("&^: %08b\n",x &^ y)
}