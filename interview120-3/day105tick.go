package main

import (
	"fmt"
	"time"
)

func statusUpdate() string { return "love" }

func main() {
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}