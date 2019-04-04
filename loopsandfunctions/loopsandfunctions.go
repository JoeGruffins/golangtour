package main

import (
	"fmt"
	"strconv"
)

var (
	enEnterString  string = `Enter number to square. "exit" to quit...`
	enAnswerString string = "Sqrt of %v is %v"
	enNumberPlz    string = `Please enter a number.`
)

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
