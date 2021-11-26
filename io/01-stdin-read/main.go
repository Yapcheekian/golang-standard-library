package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		data := make([]byte, 8)
		_, err := os.Stdin.Read(data)
		if err != nil {
			panic(err)
		}
		process(data)
	}
}

func process(data []byte) {
	fmt.Println(string(data))
}
