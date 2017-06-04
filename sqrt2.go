package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const epsilon = 1e-10
const max_iterations = 1000

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",
		float64(e))
}

type ErrNonConvergentSqrt int64

func (e ErrNonConvergentSqrt) Error() string {
	return fmt.Sprintf("Sqrt algorithm failure to converge after %v iterations",
		int64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	rand.Seed(time.Now().Unix())
	start := float64(math.Abs(rand.NormFloat64()))

	var previous float64
	var current float64
	var iteration int64

	for current = start; math.Abs(current-previous) > epsilon &&
		iteration <= max_iterations; {
		previous = current
		current = previous -
			(previous*previous-x)/(2*previous)
		iteration++
	}

	if iteration >= max_iterations {
		return x, ErrNonConvergentSqrt(iteration)
	}
	return current, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
