package main

import "fmt"

func main() {

	numbers := []float64{1, 2, 3, 4, 5}

	fmt.Println("Sum:", addUp(numbers))

	num1, num2 := tuple(numbers[0], numbers[1])

	fmt.Println("Tuple:", num1+num2)

	fmt.Printf("Params:%d\n", params(1, 2, 3))

	fmt.Printf("%d\n", funcToFunc())
}

func addUp(numbers []float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func tuple(num1 float64, num2 float64) (float64, float64) {
	return num1, num2
}

func params(nums ...float64) int {
	return int(addUp(nums))
}

func funcToFunc() int {
	num := 1
	num2 := func() int {
		num *= 2 // pointer
		return num
	}

	return num2()
}
