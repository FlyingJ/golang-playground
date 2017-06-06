package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Printf("fibonacci: sent %v to c channel\n", x)
			x, y = y, x+y
		case <-quit:
			fmt.Printf("fibonacci: received from quit channel")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("func: %v received from c channel\n", <-c)
		}
		fmt.Printf("func: sending to quit channel\n")
		quit <- 0
	}()
	fibonacci(c, quit)
}
