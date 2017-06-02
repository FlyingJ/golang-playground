package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const epsilon = 1e-10
const max_iterations = 1000 

func Sqrt(x float64) float64 {
	starting_value := float64(math.Abs(rand.NormFloat64()))

	var previous_value float64
	var current_value float64

	var iteration int64

	fmt.Printf("Initial value %v.\n", starting_value)
	fmt.Printf("%10v%25v\n", "Iteration", "Current Estimate")
	fmt.Printf("%10v%25v\n", "=========", "==================")

	for current_value = starting_value; math.Abs(current_value-previous_value) > epsilon &&
		iteration <= max_iterations; {
		previous_value = current_value
		current_value = previous_value -
			(previous_value*previous_value-x)/(2*previous_value)
		iteration++
		fmt.Printf("%10v%25v\n", iteration, current_value)
	}

	if iteration >= max_iterations {
		fmt.Printf("Bailing out after %v iterations.\n", iteration)
	} else {
		fmt.Printf("Convergence after %v iterations.\n", iteration)
	}
	return current_value
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(Sqrt(4))
}
