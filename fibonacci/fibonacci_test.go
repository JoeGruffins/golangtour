package main

import "testing"

var testoutputs = []int{
	0,
	1,
	1,
	2,
	3,
	5,
	8,
	13,
	21,
	34,
	55,
}

func TestFibonacci(t *testing.T) {
	f := fibonacci()
	for _, v := range testoutputs {
		if v != f() {
			t.Fail()
		}
	}
}
