package main

import "fmt"

func main() {
 words := []string{
	"Go", "is", "a", "high", "efficient", "language",".",
 }
 
 iw := make([]interface{}, 0, len(words))
 for _, w := range words{
	iw = append(iw, w)
 }
 fmt.Println(words...)
 fmt.Println(iw...)
}