package main

import "fmt"

type Tb struct {
	n int
}
func main() {
	ts := [2]Tb{}
	for i, t := range ts[:] {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Println(t.n)
		}
	}
	fmt.Print(ts)
}