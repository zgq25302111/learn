package main

import "fmt"

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}
/*
函数 increaseA() 是匿名返回值，返回局部变量，同时 defer 函数也会操作这个局部变量。
对于匿名返回值来说，可以假定有一个变量存储返回值，比如假定返回值变量为 anony，上面的返回语句可以拆分成以下过程：

annoy = i
i++
return

由于 i 是整型，会将值拷贝给 anony，所以 defer 语句中修改 i 值，对函数返回值不造成影响，所以返回 0 。
*/

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func sum(s int) int {
	s = 3 + 2
	return s
}

func main() {
	/*
	defer和带命名返回参数的函数一起使用时
	defer 语句定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。
	作为函数参数，则在 defer 定义时就把值传递给 defer，并被缓存起来；
	作为闭包引用(匿名函数就是闭包)的话，则会在 defer 函数真正调用时根据整个上下文确定当前的值。

	闭包是由函数及其相关引用环境组合而成的实体,即 闭包=函数+引用环境
	一般的函数都有函数名，但是匿名函数就没有。匿名函数不能独立存在，但可以直接调用或者赋值于某个变量。
	匿名函数也被称为闭包，一个闭包继承了函数声明时的作用域。在Golang中，所有的匿名函数都是闭包。

	闭包捕获的变量和常量是引用传递，不是值传递。

	https://segmentfault.com/a/1190000018169295#articleHeader4

	return xxx

	1. 返回值 = xxx
	2. 调用 defer 函数
	3. 空的 return
	*/
	x := 1
	fmt.Println(sum(x))
	fmt.Println(x)
	fmt.Println(f3())
}