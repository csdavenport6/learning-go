package main

import (
	"fmt"
	"io"
	"strings"
)

type MyLimitReader struct {
	R io.Reader
	n int64
}

func main() {
	r := strings.NewReader("bing bong")
	lr := LimitReader(r, 5)
	lr.Read([]byte("here is a string"))
	fmt.Println(lr)
}

func (lr *MyLimitReader) Read(p []byte) (n int, err error) {
	lr.R.Read(p[0:lr.n])
	return int(lr.n), nil
}

func LimitReader(r io.Reader, n int64) *MyLimitReader {
	return &MyLimitReader{r, n}
}
