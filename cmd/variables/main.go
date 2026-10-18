package main

import "fmt"

func main() {
	fmt.Println("=== Variables Lab ===")

	// Explicit type declaration
	var name string = "Go Lab"

	// Type inference
	var version = 1.22

	// Short declaration
	count := 10

	fmt.Printf("Lab: %s | Version: %.2f | Count: %d\n", name, version, count)
}
