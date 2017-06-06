package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("c has length %d and capacity %d\n", len(c), cap(c))
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	fmt.Printf("c has length %d and capacity %d\n", len(c), cap(c))
	go fibonacci(100, c)
	// fibonacci(cap(c), c)
	fmt.Printf("c has length %d and capacity %d\n", len(c), cap(c))
	for i := range c {
		fmt.Println(i)
	}
}
