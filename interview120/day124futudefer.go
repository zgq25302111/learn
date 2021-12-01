package main

import (
	"fmt"
	"errors"
)

func foo() (err error){
	defer func(){
		fmt.Println(err)
		err = errors.New("a")
	}()
	defer func(e error){
		fmt.Println(e)
		e = errors.New("b")
	}(err)
	return errors.New("c")
}

func main(){
	fmt.Println(foo())
}
