package main

import "fmt"

func main() {
	fmt.Println("=== Collections Lab ===")

	// Slices
	languages := []string{"Go", "Python", "TypeScript"}
	languages = append(languages, "Rust")
	fmt.Println("Languages:", languages)

	// Maps
	stars := map[string]int{
		"go-lab": 5,
		"demo":   2,
	}
	stars["experiments"] = 10
	fmt.Println("Stars map:", stars)
}
