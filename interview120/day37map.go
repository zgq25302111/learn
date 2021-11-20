package main

import "fmt"

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "exist", true
	}
	return "nil", false
}
func main() {
	intmap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	v, err := GetValue(intmap, 4)
	fmt.Println(v, err)
}