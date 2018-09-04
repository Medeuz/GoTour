package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	result := 1.0
	for i := 1; i <= 10; i++ {
		result = result - (result * result - x)/2*result
	}
	return result
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
