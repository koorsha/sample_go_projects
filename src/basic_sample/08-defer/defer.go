package main

import "fmt"

func main() {
	fmt.Println(saveDiv(3, 0))
	fmt.Println(saveDiv(3, 2))

}

func saveDiv(num1, num2 int) int {

	defer func() {
		recover()
	}()

	solution := num1 / num2
	return solution
}
