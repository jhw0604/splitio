package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/jhw0604/splitio"
)

func main() {
	buf := bytes.NewBufferString("hello world")
	read := splitio.New(buf, []byte(" "))
	for i := 0; i < 4; i++ {
		sub, err := read.Next()
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Println(string(sub))
		if err == io.EOF {
			break
		}
	}
}
