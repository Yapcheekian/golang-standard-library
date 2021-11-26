package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file.txt")

	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	f.Close()
	fmt.Println(string(buf))

	f, err = os.OpenFile("file2.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)

	if err != nil {
		panic(err)
	}

	io.WriteString(f, "new super dummy text")
	f.Close()
}
