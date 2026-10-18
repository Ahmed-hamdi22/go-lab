package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== Functions Lab ===")
	sum := add(10, 25)
	fmt.Printf("Sum: %d\n", sum)

	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Division result: %.2f\n", res)
	}
}
