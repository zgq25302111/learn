package main

func main() {

}
//写的可以关闭，读的关闭不了
func Stop(stop chan<- bool) {
	close(stop)
}
