package main

import (
	"fmt"
)

func main() {

	for {
		defer fmt.Println("1")

		break
	}

	fmt.Println("2")
}
