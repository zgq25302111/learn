package main

import (
	"encoding/json"
	"fmt"
)

type Peoplet struct {
	Name string `json:"name"`
}
func main() {
	js := `{
 		"name":"seekload"
 }`
	var p Peoplet
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p)
}
