package main

import "fmt"

type Param map[string]interface{}
type Show struct {
	*Param
}
func main() {
	/*s := new(Show)
	s.Param["day"] = 2*/

	s := new(Show)
	// 修复代码
	p := make(Param)
	p["day"] = 2
	s.Param = &p
	tmp := *s.Param
	fmt.Println(tmp["day"])
}
