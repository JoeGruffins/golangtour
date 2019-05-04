package main

import "testing"

const HELLO = "Hello World!"

func TestHello(t *testing.T) {
	v := hello()
	if v != HELLO {
		t.Error("Expected \"" + HELLO + "\" got \"" + v + "\"")
	}
}
