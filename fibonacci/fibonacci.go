package main

import "fmt"

func fibonacci() func() int {
	currentNum := 0
	nextNum := 1
	return func() int {
		defer func(cur, next *int) {
			tmp := *cur
			*cur = *next
			*next = *next + tmp
		}(&currentNum, &nextNum)
		return currentNum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
