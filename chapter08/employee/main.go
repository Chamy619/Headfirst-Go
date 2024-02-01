package main

import (
	"fmt"
	"headfirst-go/chapter08/magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}