package main

import (
	"fmt"
	"strconv"
)

const DELTA = 0.00000001

func nextZandDiff(x,z float64) (float64, float64) {
	newZ := z - (z*z - x) / (2*z)
	diff := newZ - z
	if diff < 0 { diff *= -1 }
	return newZ, diff
}

func Sqrt(x float64) float64 {
	z := 1.0
	newZ, diff := nextZandDiff(x,z)
	for ;diff > DELTA; z = newZ{
		newZ, diff = nextZandDiff(x,z)
	}
	return z
}

func printEnter(){
	fmt.Println("Enter number to square. \"exit\" to quit...")
}

func processInput(s string) (stop bool) {
	stop = false
	if num, err := strconv.ParseFloat(s,64); s == "exit" {
		stop = true
	} else if err != nil {
		fmt.Println("Please enter a number.")
	} else if num < 0.0 {
		fmt.Println("Only positive Numbers")
	} else {
		printEnter()
		fmt.Println(Sqrt(num))
	}
	return
}


func main() {
	printEnter()
	for{
		var input string
		fmt.Scan(&input)
		if processInput(input) { return }
	}
}

