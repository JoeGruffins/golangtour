package main

import (
	"fmt"
	"strconv"
)

const DELTA = 1e-10

var (
	enEnterString  string = `Enter number to square. "exit" to quit...`
	enAnswerString string = "Sqrt of %v is %v"
	enNumberPlz    string = `Please enter a number.`
	enNegNumString string = `Can't square negative number.`
)

type SqrtNegativeError float64

func (f SqrtNegativeError) Error() string {
	return fmt.Sprintf(enNegNumString+" %v", float64(f))
}

func nextZandDiff(x, z float64) (float64, float64) {
	newZ := z - (z*z-x)/(2*z)
	diff := newZ - z
	if diff < 0 {
		diff = -diff
	}
	return newZ, diff
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtNegativeError(x)
	}
	if x == 0 { //0 can get funky, we know its 0 so just whatevs	
	return 0, nil
	}
	z := 1.0
	timeout := 0 //some numbers will cause an oscillation, this prevents the infinite loop
	for newZ, diff := nextZandDiff(x, z); diff > DELTA && timeout < 100; z = newZ {
		newZ, diff = nextZandDiff(x, z)
	timeout++
	}
	return z, nil
}

func processInput(s string) (stop bool) {
	if num, err := strconv.ParseFloat(s, 64); s == "exit" {
		stop = true
	} else if err != nil {
		fmt.Println(enNumberPlz)
	} else {
		sqrt, err := Sqrt(num)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(enAnswerString+"\n\n"+enEnterString+"\n", num, sqrt)
	}
	return
}

func main() {
	fmt.Println(enEnterString)
	for {
		var input string
		fmt.Scan(&input)
		if processInput(input) {
			return
		}
	}
}
