package main

import "fmt"

func fibonacci() func() int {
	currentNum := 0
	nextNum := 1
	return func() int {
		defer func() {
			tmp := currentNum
			currentNum = nextNum
			nextNum += tmp
		}()
		return currentNum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
