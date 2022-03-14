package main
import "fmt"

type Foo struct {
 val int
}
func (f *Foo) Inc(inc int) {
 f.val += inc
}
func main() {
 f := &Foo{}
 f.Inc(100)
 fmt.Println(f.val)
}
