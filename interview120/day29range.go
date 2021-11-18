package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		i := i // 这⾥的 := 会᯿新声明变ᰁ，⽽不是᯿⽤
		v := v
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 3)
}
