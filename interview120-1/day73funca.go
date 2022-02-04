package main

var fa = func(i int) {
	print("x")
}
func main() {
	fa := func(i int) {
		print(i)
		if i > 0 {
			fa(i - 1)
		}
	}
	fa(10)
}
