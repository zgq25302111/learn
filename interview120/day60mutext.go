package main

import (
	"fmt"
	"sync"
	"time"
)

type datad struct {
	*sync.Mutex
}
func (d datad) test(s string) {
	d.Lock()
	defer d.Unlock()
	for i:=0;i<5 ;i++ {
		fmt.Println(s,i)
		time.Sleep(time.Second)
	}
}
func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	d := datad{new(sync.Mutex)}
	go func() {
		defer wg.Done()
		d.test("read")
	}()
	go func() {
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()
}