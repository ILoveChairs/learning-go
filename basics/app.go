// Package main requires a main function
package main

import (
	"fmt"

	"basics/conditionals"
	"basics/loops"
	"basics/print"
	"basics/variables"
)

func main() {
	fmt.Print("\nRunning Print Function!\n")
	print.Run()
	fmt.Print("\nRunning Variable Function!\n")
	variables.Run()
	fmt.Print("\nRunning Conditionals Function!\n")
	conditionals.Run()
	fmt.Print("\nRunning Loops Function!\n")
	loops.Run()
	fmt.Print("\nEnd!\n")
}
