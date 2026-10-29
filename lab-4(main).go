package main

import (
	"fmt"
	"math"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative number")
	}
	return math.Sqrt(n), nil
}

func checkInput(s string) error {
	if s == "" {
		return fmt.Errorf("input is empty")
	}
	return nil
}

func main() {
	result, err := divide(10, 2)
	fmt.Println("Division:", result, err)

	result, err = squareRoot(25)
	fmt.Println("Square root:", result, err)

	err = checkInput("")
	fmt.Println("Input check:", err)
}
