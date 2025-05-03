package models

import (
	"fmt"
)

type Staff struct {
	Person
	Address string
	Salary  int
}

func (s *Staff) SetName(name string) {
	fmt.Println("Staff SetName called")
	s.Name = name + " (Staff)"
}
