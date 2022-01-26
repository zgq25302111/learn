package main

func testa(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}
func main() {
	a, b := testa(100)
	a()
	b()
}
