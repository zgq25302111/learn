package main

import (
	"fmt"
)

func main() {
	/*var p *struct{} = nil
	fmt.Println( unsafe.Sizeof( p ) ) // 8

	var s []int = nil
	fmt.Println( unsafe.Sizeof( s ) ) // 24

	var m map[int]bool = nil
	fmt.Println( unsafe.Sizeof( m ) ) // 8

	var c chan string = nil
	fmt.Println( unsafe.Sizeof( c ) ) // 8

	var f func() = nil
	fmt.Println( unsafe.Sizeof( f ) ) // 8

	var i interface{} = nil
	fmt.Println( unsafe.Sizeof( i ) ) // 16

	type IntPtr *int
	// The underlying of type IntPtr is *int.
	var _ = IntPtr(nil) == (*int)(nil)

	// Every type in Go implements interface{} type.
	var _ = (interface{})(nil) == (*int)(nil)

	// Values of a directional channel type can be
	// converted to the bidirectional channel type
	// which has the same element type.
	var _ = (chan int)(nil) == (chan<- int)(nil)
	var _ = (chan int)(nil) == (<-chan int)(nil)

	/*var _ = ([]int)(nil) == ([]int)(nil)
	var _ = (map[string]int)(nil) == (map[string]int)(nil)
	var _ = (func())(nil) == (func())(nil)
	// The following lines compile okay.
	var _ = ([]int)(nil) == nil
	var _ = (map[string]int)(nil) == nil
	var _ = (func())(nil) == nil*/

	fmt.Println( (interface{})(nil) == (*int)(nil) ) // false
}