package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	// initialize variables
	z, d := float64(1), float64(1)

	// if delta is greater than 1e-15, seek closer approximation
	for d > 1e-15 {
		z0 := z
		z = z - (z*z-x)/(2*z)
		d = math.Abs(z - z0)
	}
	return z

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
