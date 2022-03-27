package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

func main() {
	var mu MyMutex
	mu.Lock()
	var mu1 = mu
	mu.count++
	mu.Unlock()
	mu1.Unlock()
	mu1.count++
	fmt.Println(mu.count, mu1.count)
}
