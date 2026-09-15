package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p Person) ReadDatafromUserinput() Person {
	fmt.Print("Enter name: ")
	fmt.Scan(&p.Name)
	fmt.Print("Enter age: ")
	fmt.Scan(&p.Age)
	fmt.Print("Enter job: ")
	fmt.Scan(&p.Job)
	fmt.Print("Enter salary: ")
	fmt.Scan(&p.Salary)
	return p
}

func (p Person) PrintallData() {
	fmt.Print("Name: ", p.Name)
	fmt.Print("\nAge: ", p.Age)
	fmt.Print("\nJob: ", p.Job)
	fmt.Print("\nSalary: ", p.Salary)
}

func main() {
	var p2 Person
	var p1 Person
	p1 = p1.ReadDatafromUserinput()
	p2 = p2.ReadDatafromUserinput()
	p1.PrintallData()
	p2.PrintallData()
}
