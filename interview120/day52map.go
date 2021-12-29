package main

import "fmt"

func main() {
	x := map[string]string{"one":"1", "two":"2", "three":"3"}
	if _,ok := x["two"]; !ok{
		fmt.Println("blank")
	}
}
