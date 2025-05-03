package main

import (
	"fmt"
	"reflect"
)
import (
	"go-learning/utils"
)
import (
	"go-learning/models"
)

var name string = "Go"

func main() {
	age := 10
	age_1 := 11
	var age_2 int = utils.Add(age_1, age)
	fmt.Println("Hello, world!", age, age, reflect.TypeOf(age).Name())
	fmt.Println("Add utils: ", age_2)

	s := models.Staff{
		Person:  models.Person{Name: "Thanh", Age: 25}, // lỗi nếu không import đúng
		Address: "Vietnam",
		Salary:  1000,
	}

	fmt.Println("before SetName", s)
	s.SetName(s.Name)
	fmt.Println("after SetName", s.Name)
}
