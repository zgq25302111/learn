//第2天
//https://mp.weixin.qq.com/s/IWqYVhL5tLe576oP2KxAGA
package main

import "fmt"

func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)
	for key,val := range slice {
		m[key] = &val
		fmt.Println(key,"->",*m[key])
	}
}

/*type Test struct {
	name string
}
func (this *Test) Point(){
	fmt.Println(this.name)
}
func main() {
	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}
	for _,t := range ts {
		//fmt.Println(reflect.TypeOf(t))
		t.Point()
	}
}*/

/*
func maint() {
	v := []int{1, 2, 3, 4, 5}
	for _,x := range v {
		v = append(v, x)
		fmt.Println(v)
	}
names := []string{"John", "Bob", "Claire", "Nik"}

	for i, name := range names {
		fmt.Println("Element at index", i, "=", name)
	}

	names = []string{"John", "Bob", "Claire", "Nik"}
	for _, name := range names {
		name = strings.ToUpper(name)
	}
	fmt.Println(names)

	for i := range names {
		names[i] = strings.ToUpper(names[i])
	}
	fmt.Println(names)

	//Merge two slices
	b := append(b, c...)
	test := append([]int{10,20}, []int{30,40,50}...)

	names := []string{"John", "Bob", "Claire", "Nik"}
	for i, name := range names {
    	if name == "Claire" {
        	fmt.Println("Claire found at index", i)
	}

	//delete element
	a = append(a[:i], a[i+1:]...)
	//Put an element at index
	// objective put "C" at index 2
// index:            0    1    2    3    4
letters := []string{"A", "B", "D", "E", "F"}

// 1) add an element to the end of the slice
letters = append(letters, "")

// 2) copy letters[i:] to letters[i+1:]
copy(letters[3:], letters[2:])

// 3) set "C" at index 2
letters[2] = "C"

//Generalization
s = append(s, 0)
copy(s[i+1:], s[i:])
s[i] = x

}
*/