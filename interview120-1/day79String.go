package main

import "fmt"

type ConfigOnea struct {
	Daemon string
}
func (c *ConfigOnea) String() string {
	return fmt.Sprintf("print: %v", c)
}
func main() {
	c := &ConfigOnea{}
	c.String()
}
