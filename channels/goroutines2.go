package main

import (
	"fmt"
	"math/rand"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		var dt time.Duration
		switch s {
		case "world":
			dt = time.Duration(100+rand.Int63n(400)) * time.Millisecond
		default:
			dt = 100 * time.Millisecond
		}
		time.Sleep(dt)
		fmt.Printf("%d: %s\n", i, s)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	/*
		x := rand.Int63n(100)
		fmt.Printf("x has type %T\n", x)

		dt := time.Duration(x) * time.Millisecond
		fmt.Printf("dt has type %T\n", dt)
	*/
	go say("world")
	say("hello")
}
