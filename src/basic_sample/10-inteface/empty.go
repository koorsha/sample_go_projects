package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	switch x.(type) {
	case int:
		fmt.Println("Integer!")
	case string:
		fmt.Println("String")
	}
}
