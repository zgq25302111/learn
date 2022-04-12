package main
import "fmt"

func main() {
 a := []int{1, 2, 3, 4}
 b := variadic(a...)
 b[0], b[1] = b[1], b[0]
 fmt.Println(a)
}
func variadic(ints ...int) []int {
 return ints
}