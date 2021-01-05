package main

import "fmt"

func main() {

	// first way to get value from memeory address
	s1 := "My String 1"
	p1 := &s1        // get address
	fmt.Println(*p1) // get value via point to address

	// 2 way
	s2 := "My String 2"
	p2 := new(string) // create empty
	p2 = &s2          // assign to mem address
	fmt.Println(*p2)  // get value via point to address

	// 3 way
	s3 := "My String 3"
	var p3 *string // is yet nil
	p3 = &s3       // assign to mem address
	fmt.Println(*p3)

	x := 0 // init empty value
	ptr(&x) // pass mem address to pointer
	fmt.Println(x)
	fmt.Println(&x)

	// or
	num := new(int) // create a nil address in mem
	ptr(num) // pass mem address
	fmt.Println(*num) // get value via point to address
	fmt.Println(&num)
}

func ptr(x *int) {
	*x = 2
}
