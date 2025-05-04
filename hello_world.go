package main

import (
	"fmt"
	"reflect"
)
import (
	"go-learning/utils"
)
import (
	"go-learning/database"
	"go-learning/models"
)
import (
	"go-learning/simple"
)

func main() {
	age := 10
	fmt.Println("1. Hello, world!", age, age, reflect.TypeOf(age).Name())
	age_1 := 11
	var age_2 int = utils.Add(age_1, age)

	fmt.Println("2. Simple function: ", age_2)

	s := models.Staff{
		Person:  models.Person{Name: "Thanh", Age: 25}, // lỗi nếu không import đúng
		Address: "Vietnam",
		Salary:  1000,
	}

	fmt.Println("3. Simple Struct SetName", s)
	s.SetName(s.Name)
	fmt.Println("4. Simple Struct SetName", s.Name)

	fmt.Println("Simple for loop")
	simple.ForLoop(5)
	switchCaseReturn := simple.SwitchCase(10)
	fmt.Println("5. Switch Case Return", switchCaseReturn)
	switchCaseReturn = simple.SwitchCase(1)
	fmt.Println("5. Switch Case Return", switchCaseReturn)

	database.ConnectDB()

	database.MigrateAll(database.Db, "up")
}
