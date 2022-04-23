package main

import (
	"fmt"
	"errors"
)

func foo() (err error){

	/*defer 先进后出 b上边打印 那个e是nil 因为 defer函数没有参数的 
	然后是a上边打印 通过foo返回参数捕获c 打印完c后又改变err值 所以最后打印a
	*/
	
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
	//fmt.Println(foo())
	foo()
}
