package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer([]byte{})

	if err := binary.Write(buf, binary.BigEndian, 1.33); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
	fmt.Println(1.33)

	var num float64
	binary.Read(buf, binary.BigEndian, &num)
	fmt.Println(num)
}
