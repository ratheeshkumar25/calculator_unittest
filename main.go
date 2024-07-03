package main

import (
	"fmt"

	"github.com/ratheeshkumar25/calculator"
)

func main() {
	result := calculator.Add(2,5)
	result1 := calculator.Add(2,3)
	result2 := calculator.Sub(9,2)
	result3 := calculator.Multipication(5,5)
	result4 := calculator.Division(25,5)

	fmt.Printf("5 + 2 = %d\n", result)
	fmt.Printf("2 + 3 = %d\n", result1)
	fmt.Printf("9 - 2 = %d\n", result2)
	fmt.Printf("5 * 5 = %d\n", result3)
	fmt.Printf("25 / 2 = %d\n", result4)
}