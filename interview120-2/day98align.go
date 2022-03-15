package main

import "fmt"

func test(i int) (ret int) {
 ret = i * 2
 if ret > 10 {
 //ret := 10
 return
 }
 return
}
func main() {
 result := test(10)
 fmt.Println(result)
}