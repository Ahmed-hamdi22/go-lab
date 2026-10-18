package main

import "fmt"

func main() {
	fmt.Println("=== Conditions Lab ===")

	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}

	// Switch example
	day := "Monday"
	switch day {
	case "Friday", "Saturday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Workday!")
	}
}
