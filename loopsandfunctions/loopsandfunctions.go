package main

import (
	"fmt"
	"strconv"
)

const DELTA = 0.00000001

var (
	enEnterString  string = `Enter number to square. "exit" to quit...`
	enNegNumString string = `Can't square negative number.`
	enAnswerString string = "Sqrt of %v is %v"
	enNumberPlz    string = `Please enter a number.`
)

type SqrtNegativeError float64

func (f SqrtNegativeError) Error() string {
	return fmt.Sprintf(enNegNumString+" %v", float64(f))
}

func nextZandDiff(x, z float64) (float64, float64) {
	newZ := z - (z*z-x)/(2*z)
	diff := newZ - z
	if diff < 0 {
		diff *= -1
	}
	return newZ, diff
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtNegativeError(x)
	}
	z := 1.0
	newZ, diff := nextZandDiff(x, z)
	for ; diff > DELTA; z = newZ {
		newZ, diff = nextZandDiff(x, z)
	}
	return z, nil
}

func processInput(s string) (stop bool) {
	stop = false
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
		fmt.Printf(enAnswerString+"\n\n", num, sqrt)
		fmt.Println(enEnterString)
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
