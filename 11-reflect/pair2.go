package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// cmd: pair<type:*os.File, value:"cmd">
	cmd, err := os.OpenFile("cmd", os.O_RDWR, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	//r: pair<type: , value: >
	var r io.Reader
	//r: pair<type:*os.File, value:"cmd">
	r = cmd

	var w io.Writer
	w = r.(io.Writer)

	write, err := w.Write([]byte("this is a test.\n"))
	if err != nil {
		return
	}
	fmt.Println("write =", write)
}
