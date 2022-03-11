package main

import "fmt"

func get() []byte {
 raw := make([]byte, 10000)
 fmt.Println(len(raw), cap(raw), &raw[0])
 res := make([]byte, 3)
 copy(res, raw[:3])
 return res
}

func main() {
 data := get()
 fmt.Println(len(data), cap(data), &data[0])
}
