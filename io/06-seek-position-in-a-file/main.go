package main

import (
	"errors"
	"fmt"
	"os"
)

const lineLength = 25

func main() {
	f, err := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	c, _ := readRecord(0, "first", f)
	fmt.Println(c)

	if err := writeRecord(1, "id", "333", f); err != nil {
		panic(err)
	}

	c, _ = readLine(2, f)
	fmt.Println(c)
}

func readLine(line int, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.Seek(int64(line*lineLength), 0)
	_, err := f.Read(lineBuffer)
	return string(lineBuffer), err
}

func writeRecord(line int, column string, dataStr string, f *os.File) error {
	definedLength := 10
	position := int64(line * lineLength)
	switch column {
	case "id":
		definedLength = 4
	case "first":
		position += 4
	case "last":
		position += 14
	default:
		return errors.New("column not defined")
	}

	data := make([]byte, definedLength)
	for i := range data {
		data[i] = '.'
	}
	copy(data, []byte(dataStr))
	_, err := f.WriteAt(data, position)
	return err
}

func readRecord(line int, column string, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.ReadAt(lineBuffer, int64(line*lineLength))
	var retVal string
	switch column {
	case "id":
		return string(lineBuffer[:3]), nil
	case "first":
		return string(lineBuffer[4:13]), nil
	case "last":
		return string(lineBuffer[14:]), nil
	}
	return retVal, errors.New("Column not defined")
}
