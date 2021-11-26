package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "Hello World from stdout\n")
	io.WriteString(os.Stderr, "Hello World from stderr\n")
	os.Stdout.Write([]byte("Hello World from stdout\n"))
	fmt.Fprintln(os.Stderr, "Hello World from stderr")
}
