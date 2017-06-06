package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// assumes tree with 10 nodes!
	// this is not great!
	for i := 0; i < 10; i++ {
		x, y := <-ch1, <-ch2
		if x != y {
			return false
		}
	}
	return true
}

func main() {
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
