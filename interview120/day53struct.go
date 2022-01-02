package main

import "fmt"

type Tx struct {
	n int
}

func (t *Tx) Set(n int) {
	t.n = n
}
func getT() Tx {
	return Tx{}
}
func main() {
	t := getT()
	t.Set(2)
	fmt.Println(t.n)
}