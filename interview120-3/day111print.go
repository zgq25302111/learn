package main

func main() {
	var val int
	a := &val
	println(a)
	f(10000)
	b := &val
	println(a) // a b 的值相同
	println(b)
	println(a == b) // A
}
func f(i int) {
	if i--; i == 0 {
		return
	}
	f(i)
}
