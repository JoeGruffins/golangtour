package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

const HELLO = "Hello World!\n"

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

func TestHello(t *testing.T) {
	v := captureOutput(func() {
		hello()
	})
	if v != HELLO {
		t.Error("\nExpected:\n" + HELLO + "got:\n" + v)
	}
}
