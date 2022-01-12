package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func main() {
	var w WordCounter
	w.Write([]byte("word word word word word"))
	fmt.Println("word counter: ", w)

	var l LineCounter
	l.Write([]byte("line\nline\nline\nline\nline\n"))
	fmt.Println("line counter: ", l)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanWords)
	*w += WordCounter(count)
	return count, nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanLines)
	*l += LineCounter(count)
	return count, nil
}

func retCount(p []byte, fn bufio.SplitFunc) (count int) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(fn)
	count = 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
	}
	return
}
