package main

import "fmt"

type BasePerson struct {
	First string
	Last  string
}

type Employee struct {
	BasePerson
	Salary int
	Boss   *Manager
}

type Manager struct {
	Employee
}

type Person interface {
	GetName() string
}

func ToString(p Person) {
	fmt.Printf("Hello, %s\n" , p.GetName())
}

func (e Employee) GetName() string {
	return e.First
}

func (m Manager) GetName() string {
	return m.First
}

func main() {
	m := &Manager{
		Employee{
			BasePerson: BasePerson{
				First: "Koorsha",
				Last:  "Shirazi",
			},
			Salary: 1000,
			Boss:   nil,
		},
	}
	e := &Employee{
		BasePerson: BasePerson{
			First: "Armin",
			Last:  "Naderi",
		},

		Salary: 2000,
		Boss:   m,
	}

	ToString(m)
	ToString(e)

	people :=[]Person{Person(m),Person(e)}
	for _, p:= range(people) {
		ToString(p)
	}
}
