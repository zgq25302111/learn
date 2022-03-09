package main

import (
	"fmt"
	"time"
)

func Ac() int {
	time.Sleep(100 * time.Millisecond)
	return 1
}
func Bc() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}
func main() {
	ch := make(chan int, 1)
	go func() {
		select {
		case ch <- Ac():
		case ch <- Bc():
		default:
			ch <- 3
		}
	}()
	fmt.Println(<-ch)
}