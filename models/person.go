package models

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SetName(name string) {
	fmt.Println("Person SetName called")
	p.Name = name
}
