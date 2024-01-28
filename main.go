package main

import (
	"fmt"

	"github.com/kunchalavikram1427/go-calculator/calculator"
	"github.com/kunchalavikram1427/go-calculator/utils"
)

func main() {
	fmt.Println("Hello from myapp!")

	// Using functions from the calculator package
	resultAdd := calculator.Add(5, 3)
	resultSubtract := calculator.Subtract(8, 4)

	fmt.Printf("Addition result: %d\n", resultAdd)
	fmt.Printf("Subtraction result: %d\n", resultSubtract)

	// Using a function from the utils package
	isValid := utils.Validate("some input")
	fmt.Printf("Validation result: %v\n", isValid)
}
