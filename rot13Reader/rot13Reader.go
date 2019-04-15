package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func isUpper(c *byte) bool {
	return *c >= 'A' && *c <= 'Z'
}

func isLower(c *byte) bool {
	return *c >= 'a' && *c <= 'z'
}

func shift(b *byte) {
	if isLower(b) || isUpper(b) {
		var c byte
		if isUpper(b) {
			c = 'A'
		} else {
			c = 'a'
		}
		shiftBits := (*b - c + 13) % 26
		*b = c + shiftBits
	}
	return
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	for i := range b {
		shift(&b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
