package main

import "fmt"

/*func main() {
	var x = nil
	_ = x
}
*/

func main() {
	var x interface{} = nil
	_ = x
	fmt.Println(x)
}