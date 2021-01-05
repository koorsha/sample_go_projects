package main

import "fmt"

func main() {
	var arr [5]float64
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	arr[3] = 400
	arr[4] = 500

	for i, value := range arr {
		fmt.Println(value, i)
	}

	arr2 := [5]float64{1, 2, 3, 4, 5}
	for _, value := range arr2 {
		fmt.Println(value)
	}

}
