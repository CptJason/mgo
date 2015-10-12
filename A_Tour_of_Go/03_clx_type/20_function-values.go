package main

import (
	"fmt"
	"math"
)

func main() {
	// Function values
	// Functions are values too.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}
