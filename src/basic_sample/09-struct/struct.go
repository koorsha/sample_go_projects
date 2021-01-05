package main

import (
	"fmt"
	"time"
)

// Person Type
type Person struct {
	First string
	Last  string
	Dob   time.Time
}

// CreatePerson return new instance of person
// CreatePerson is a extension method and public
func CreatePerson(f, l string, y, m, d int) *Person {
	return &Person{
		First: f,
		Last:  l,
		Dob:   time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local),
	}
}

// FullNme return first and last of person
func (person Person) FullNme() string {
	return person.First + "\t" + person.Last
}

// GetAge
func (person Person) GetAge() int {
	d := time.Since(person.Dob)
	return int(d.Hours() / 24 / 365)
}

func main() {
	p1 := CreatePerson("koorsha", "Shirazi", 1987, 2, 6)
	fmt.Printf("Hello, %s\n", p1.FullNme())
	fmt.Printf("Your age is, %d\n", p1.GetAge())
}
