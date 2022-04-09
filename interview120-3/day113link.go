package main

import "fmt"

func link(p ...interface{}) {
	fmt.Println(p)
}
func main() {
	link("seek", 1, 2, 3, 4) // 输出 [seek 1 2 3 4]
	a := []int{1, 2, 3, 4}
	link("seek", a) // 输出 [seek [1 2 3 4]]
	tmplink := make([]interface{}, 0, len(a)+1)
	tmplink = append(tmplink, "seek")
	for _, ii := range a {
		tmplink = append(tmplink, ii)
	}
	link(tmplink...) // 输出 [seek 1 2 3 4]
}
