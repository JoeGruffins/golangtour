package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"testing"
)

type sqrtTestTrio struct {
	input     float64
	outputNum float64
	outputErr error
}

type inputTestTrio struct {
	input        string
	outputBool   bool
	outputStdout string
}

var sqrtTests = []sqrtTestTrio{
	{4, 2, nil},
	{35, math.Sqrt(35), nil},
	{0.1, math.Sqrt(0.1), nil},
	{5.4e13, math.Sqrt(5.4e13), nil},
	{0, 0, nil},
	{-1, 0, SqrtNegativeError(-1)},
	{outputNum: 0, outputErr: nil},
}

var inputTests = []inputTestTrio{
	{"exit", true, ""},
	{"eight", false, fmt.Sprintf(enNumberPlz + "\n")},
	{"", false, fmt.Sprintf(enNumberPlz + "\n")},
	{"quit", false, fmt.Sprintf(enNumberPlz + "\n")},
	{`!"#$%&'()@<>`, false, fmt.Sprintf(enNumberPlz + "\n")},
	{"54ee12", false, fmt.Sprintf(enNumberPlz + "\n")},
	{"5e12", false, fmt.Sprintf(enAnswerString+"\n\n"+enEnterString+"\n", 5e12, sqrtNoError(5e12))},
	{"7", false, fmt.Sprintf(enAnswerString+"\n\n"+enEnterString+"\n", 7, sqrtNoError(7))},
	{"0", false, fmt.Sprintf(enAnswerString+"\n\n"+enEnterString+"\n", 0, 0)},
	{"-3", false, fmt.Sprintf(SqrtNegativeError(-3).Error() + "\n")},
}

//helper func to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = old
	return <-outC
}

//helper function to ignore error
func sqrtNoError(x float64) float64 {
	v, _ := Sqrt(x)
	return v
}

func TestSqrt(t *testing.T) {
	for _, trio := range sqrtTests {
		v, err := Sqrt(trio.input)
		if v != trio.outputNum || err != trio.outputErr {
			t.Error("for", trio.input, "expected", trio.outputNum, trio.outputErr, "got", v, err)
		}
	}
}

func TestInput(t *testing.T) {
	for _, trio := range inputTests {
		var v bool
		stdout := captureOutput(func() {
			v = processInput(trio.input)
		})
		if v != trio.outputBool || stdout != trio.outputStdout {
			t.Error("for", trio.input, "expected", trio.outputBool, trio.outputStdout, "got", v, stdout)
		}
	}
}
