package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	buf := bytes.NewBuffer([]byte{})

	f, _ := os.OpenFile("file.txt", os.O_RDWR, os.ModePerm)

	w := io.MultiWriter(f, buf)

	io.WriteString(w, "hello")

	fmt.Println(buf.String())
}
