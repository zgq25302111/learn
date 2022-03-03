package main

type Usera struct {
	Name string
}

func (u *Usera) SetName(name string) {
	u.Name = name
}

type Employeea struct {
	Usera
	Title      string
}

func main(){
	employee := new(Employeea)
	employee.SetName("Jack")
	// error employee.SetName undefined (type *Employee has no field or method SetName)
}