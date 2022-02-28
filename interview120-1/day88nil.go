package main

import (
	"fmt"
)

func Fooa() (err error) {
	//var err *os.PathError = nil
	// …
	return err
}
func main() {
	err := Fooa()
	fmt.Println(err)
	fmt.Println(err == nil)
	fmt.Printf("%#v\n",err)
}