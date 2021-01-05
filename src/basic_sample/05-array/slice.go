package main

import "fmt"

func main() {

	num := []int{1, 2, 5, 3, 9, 8, 7, 6}
	num2 := num[1:8]

	fmt.Println(num2[1:])
	fmt.Println(num2[:3])

	num3 := make([]int, 5, 10)

	copy(num3, num)
	fmt.Println(num3[0:])

	num4 := append(num3, 10)
	fmt.Println(num4[0:])

	for _, value := range num2 {
		fmt.Println(value)
	}
}
