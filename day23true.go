package main

import "fmt"

func main() {
	/*
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
	*/

	//以上代码等价于

	{
		a := 1
		if false {

		} else {
			{
				b := 2
				if false {

				} else {
					fmt.Println(a, b)
				}
			}
		}
	}

}
