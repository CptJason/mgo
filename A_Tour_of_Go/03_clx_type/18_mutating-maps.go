package main

import "fmt"

func main() {
	m := make(map[string]int)

	// Insert or update an element in map m:
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// Insert or update an element in map m:
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Retrieve an element:
	elem := m["Answer"]

	// Test that a key is present with a two-value assignment:
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok, elem)
}
