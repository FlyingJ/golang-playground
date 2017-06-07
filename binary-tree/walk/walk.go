package walk

import "golang.org/x/tour/tree"

// Walk begins the process of walking through a tree given root
// node root
// walks the tree t sending all values
// from the tree to the channel ch.
func Walk(root *tree.Tree, ch chan int) {
	if root.Left != nil {
		subTreeWalk(root.Left, ch)
	}
	ch <- root.Value
	if root.Right != nil {
		subTreeWalk(root.Right, ch)
	}
	close(ch)
}

// subTreeWalk walks subtree t sending values
// from the tree to the channel ch
// the difference is that subTreeWalk does not
// close the value collection channel
func subTreeWalk(child *tree.Tree, ch chan int) {
	if child.Left != nil {
		subTreeWalk(child.Left, ch)
	}
	ch <- child.Value
	if child.Right != nil {
		subTreeWalk(child.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)

	ch2 := make(chan int)
	go Walk(t2, ch2)
	
	for {
		x, ch1ok := <-ch1
		y, ch2ok := <-ch2
		if (!ch1ok && ch2ok) || (ch1ok && !ch2ok) {
			return false
		}
		if x != y {
			return false
		}
		return true
	}
}
