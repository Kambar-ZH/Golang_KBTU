package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
	len, err := rot13.r.Read(b)
	if (err != nil) {
		return 0, err
	}
	for i := 0; i < len; i++ {
		if (b[i] >= 'A' && b[i] <= 'Z') {
			b[i] += 13
			if (b[i] > 'Z') {
				b[i] -= 26
			}
		} else if (b[i] >= 'a' && b[i] <= 'z') {
			b[i] += 13
			if (b[i] > 'z') {
				b[i] -= 26
			}
		}
	}
	return len, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}