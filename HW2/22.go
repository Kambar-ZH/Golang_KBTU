package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	len := len(b)
	for i := 0; i < len; i++ {
		b[i] = 'A'
	}
	return len, nil
}

func main() {
	reader.Validate(MyReader{})
}