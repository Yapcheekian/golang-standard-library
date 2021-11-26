package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {
	pReader, pWriter := io.Pipe()

	cmd := exec.Command("ls", "-l")
	cmd.Stdout = pWriter

	go func() {
		defer pReader.Close()
		io.Copy(os.Stdout, pReader)
	}()

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
