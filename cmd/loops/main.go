package main

import "fmt"

func main() {
	fmt.Println("=== Loops Lab ===")

	// Standard for loop
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration: %d\n", i)
	}

	// While-style loop
	count := 3
	for count > 0 {
		fmt.Printf("Countdown: %d\n", count)
		count--
	}
}
