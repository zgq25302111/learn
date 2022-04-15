package main

import (
	"fmt"
	"errors"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}
func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}
func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}
