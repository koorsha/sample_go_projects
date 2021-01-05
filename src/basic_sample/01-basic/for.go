package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
		for {

			time.Sleep(1 * time.Second)
			break
		}
	}

}
