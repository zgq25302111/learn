package main

import (
	"fmt"
)

func main(){
	s := []int{1, 2, 3, 4}
	s = Append(s)
	fmt.Println(s)
	Add(s)
	fmt.Println(s)
}

func Append(s []int)(t []int){
	s = append(s, 5)
	return s
}

func Add(s []int){
	for i:= range s{
		s[i]=s[i]+5
	}
}

/*
func main(){
	s := []int{1, 2, 3, 4}
	Append(&s)
	fmt.Println(s)
	Add(s)
	fmt.Println(s)
}

func Append(s *[]int){
	*s = append(*s, 5)
}
*/