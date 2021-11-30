package main

import (
  "fmt"
  "time"
)

func writeToChannel(c chan int, x int){
  fmt.Println(x)
  c <- x
  close(c)
  y := <-c
  y++
  fmt.Println(y)
}

func main(){
  c := make(chan int)
  go writeToChannel(c, 10)
  time.Sleep(1 * time.Second)
}
