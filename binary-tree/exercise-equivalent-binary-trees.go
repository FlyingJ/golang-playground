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
//func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)
	select {
	case i := <-ch:
		fmt.Println(i)
	default:
		fmt.Println("done")
	}

}

