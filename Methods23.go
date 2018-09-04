package main

import (
	"io"
	"os"
	"strings"
)

type Rot13Reader struct {
	r io.Reader
}

func (reader Rot13Reader) Read(b []byte) (int, error) {
	length, error := reader.r.Read(b)

	if (error != io.EOF) {
		for i := 0; i < length; i++ {
			if b[i] < 'N' || (b[i] >= 'a' && b[i] < 'n') {
					b[i] += 13
				} else {
					b[i] -= 13
				}
		}
	}

	return length, io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
