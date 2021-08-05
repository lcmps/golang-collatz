package main

import (
	"fmt"
	"os"
)

func main() {
	Collatz(9)
}

func Collatz(integer int) {
	if integer == 1 {
		os.Exit(0)
	}

	if integer%2 == 0 {
		integer = integer / 2
		fmt.Println(integer)
		Collatz(integer)
	}

	integer = (3 * integer) + 1
	fmt.Println(integer)
	Collatz(integer)
}
