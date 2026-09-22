package main

import "fmt"

func birthday(age *int) int {
	*age = *age + 1
	return *age
}

type Person struct {
	Name string
	Age  int
}

func main() {
	var x int = 10
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of x:", x)

	birthday(&x)
	fmt.Println("Age Before Birthday:", x)
	fmt.Println("Age after birthday:", birthday(&x))

	p := new(Person)
	p.Name = "Muskan"
	p.Age = 22
	fmt.Println("Person before modification:", *p)
	p.Age = 23
	fmt.Println("Person after modification:", *p)

}
