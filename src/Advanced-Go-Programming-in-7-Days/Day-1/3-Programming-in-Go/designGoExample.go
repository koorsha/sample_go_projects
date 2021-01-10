package main

import "fmt"

func between(from, to int) []int {
	var result []int

	if from > to {
		result = []int{}
	} else {
		for i := from; i < to; i++ {
			result = append(result, i)
		}
	}

	return result
}

// WhiteSpace use Case list
func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

// ExitWithBreak A break statement terminates execution of the innermost for, switch, or
// select statement.
func ExitWithBreak() {
Loop:
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ': // skip space
			break
		case '\n': // break at newline
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}

// Fallthrough A fallthrough statement transfers control to the next case.
func Fallthrough() {
	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}

func main() {

	// A switch statement runs the first case equal to the condition expression.
	// The cases are evaluated from top to bottom, stopping when a case succeeds.
	// If no case matches and there is a default case, its statements are executed.
	// https://yourbasic.org/golang/switch-statement/
	for i := range between(0, 10) {
		switch i % 5 {
		case 1:
			fmt.Println("fizz")
		case 2:
			fmt.Println("bazz")
		case 3:
			fmt.Println("gizz")
			fallthrough
		default:
			fmt.Println("fizzbazz")
		}
	}
}
