package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := bytes.Buffer{}
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		buf.WriteString(sc.Text())
	}

	fmt.Println(buf.String())

	b, _ := ioutil.ReadFile("file.txt")

	fmt.Println(string(b))
}
