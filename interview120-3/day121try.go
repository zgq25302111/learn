package main

func main(){
	println(intFunc(7))
}

func intFunc(myint int) int {
	
	x := 5
	
	if myint == 1{
		x = 1
		return x
	}
	
	if myint == 2{
		x = 2
		return x
	}
	
	if myint == 3{
		x = 3
		return x
	}
	
	if myint == 4{
		x = 4
		return x
	}
	
	return x
}