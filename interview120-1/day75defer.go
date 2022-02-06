package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	println(string(b))
}