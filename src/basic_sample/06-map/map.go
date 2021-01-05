package main

import "fmt"

func main() {

	person := make(map[string]string)
	person["Name"] = "Koorsha"
	person["Family"] = "Shirazi"
	person["Age"] = "32"

	fmt.Println(person["Name"])
	fmt.Println(len(person))

	delete(person, "Age")
	fmt.Println(len(person))
}
