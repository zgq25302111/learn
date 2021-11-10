package main

import "fmt"

type Math struct {
	x, y int
}
var mm = map[string]*Math{
	"foo": &Math{2, 3},
}
func main() {
	mm["foo"].x = 4
	fmt.Println(mm["foo"].x)
	fmt.Printf("%#v", mm["foo"]) // %#v 格式化输出详细信息
}