package main

import "fmt"

func increment(val *int) {
	*val = *val + 1
}

func main() {
	fmt.Println("=== Pointers Lab ===")

	num := 41
	fmt.Println("Before increment:", num)

	increment(&num)
	fmt.Println("After increment via pointer:", num)
}
