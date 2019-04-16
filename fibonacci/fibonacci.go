package main

import "fmt"

func fibonacci() func() int {
	currentNum, nextNum := 0, 1
	return func() int {
		defer func() {
			currentNum, nextNum = nextNum, currentNum+nextNum
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
