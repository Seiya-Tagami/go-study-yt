package main

import "fmt"

// Non-generic
func sumOfInt(input []int) int {
	inc := 0
	for _, val := range input {
		inc += val
	}

	return inc
}

// Generic
type Number interface {
	int64 | float64
}

func sumOf[T Number](input []T) T {
	var inc T
	for _, val := range input {
		inc += val
	}
	return inc
}

func main() {
	fmt.Printf("%d", sumOf([]int64{0,1,2}))
	fmt.Printf("%0.2f", sumOf([]float64{0.0,1.0,2.4}))
}