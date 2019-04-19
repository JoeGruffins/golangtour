package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *tree.Tree) chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func Same(t1, t2 *tree.Tree) bool {
	w1 := Walker(t1)
	w2 := Walker(t2)
	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2

		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

func main() {
	t1 := tree.New(2)
	t2 := tree.New(2)
	fmt.Printf("same: %v", Same(t1, t2))
}
