package main

import "fmt"

type Map map[string]string

func (m Map) Set(key string, value string) {
	m[key] = value
}
func main() {
	m := make(Map)
	m.Set("A", "One")
	fmt.Println(m["A"])
}
