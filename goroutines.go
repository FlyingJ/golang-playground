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
	go say("world")
	say("hello")
}
