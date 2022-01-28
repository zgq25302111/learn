package main

func testb() []func() {
	var funs []func()
	for i := 0; i < 3; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	return funs
}
func main() {
	funs := testb()
	for _, f := range funs {
		f()
	}
}